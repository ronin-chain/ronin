// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package miner

import (
	"bytes"
	"errors"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	mapset "github.com/deckarep/golang-set"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/consortium"
	"github.com/ethereum/go-ethereum/consensus/misc"
	"github.com/ethereum/go-ethereum/consensus/misc/eip1559"
	"github.com/ethereum/go-ethereum/consensus/misc/eip4844"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/txpool"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/trie"
	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/holiman/uint256"
)

const (
	// resultQueueSize is the size of channel listening to sealing result.
	resultQueueSize = 10

	// txChanSize is the size of channel listening to NewTxsEvent.
	// The number is referenced from the size of tx pool.
	txChanSize = 4096

	// chainHeadChanSize is the size of channel listening to ChainHeadEvent.
	chainHeadChanSize = 10

	// chainSideChanSize is the size of channel listening to ChainSideEvent.
	chainSideChanSize = 10

	// resubmitAdjustChanSize is the size of resubmitting interval adjustment channel.
	resubmitAdjustChanSize = 10

	// miningLogAtDepth is the number of confirmations before logging successful mining.
	miningLogAtDepth = 7

	// minRecommitInterval is the minimal time interval to recreate the mining block with
	// any newly arrived transactions.
	minRecommitInterval = 1 * time.Second

	// maxRecommitInterval is the maximum time interval to recreate the mining block with
	// any newly arrived transactions.
	maxRecommitInterval = 15 * time.Second

	// intervalAdjustRatio is the impact a single interval adjustment has on sealing work
	// resubmitting interval.
	intervalAdjustRatio = 0.1

	// intervalAdjustBias is applied during the new resubmit interval calculation in favor of
	// increasing upper limit or decreasing lower limit so that the limit can be reachable.
	intervalAdjustBias = 200 * 1000.0 * 1000.0

	// staleThreshold is the maximum depth of the acceptable stale block.
	staleThreshold = 7

	// recentMinedCacheLimit is the maximum blocks stored in recentMinedBlocks for avoiding
	// double sign a block
	recentMinedCacheLimit = 20

	// maxBlockSize is the maximum size of block, this is used to abort transaction commit
	// to block when the estimated block size is larger than this threshold
	maxBlockSize = 10 * 1024 * 1024 // 10MB
)

// environment is the worker's current environment and holds all of the current state information.
type environment struct {
	signer types.Signer

	state              *state.StateDB // apply state changes here
	ancestors          mapset.Set     // ancestor set (used for checking uncle parent validity)
	family             mapset.Set     // family set (used for checking uncle invalidity)
	uncles             mapset.Set     // uncle set
	tcount             int            // tx count in cycle
	gasPool            *core.GasPool  // available gas used to pack transactions
	estimatedBlockSize uint64

	header   *types.Header
	txs      []*types.Transaction
	receipts []*types.Receipt

	sidecars []*types.BlobTxSidecar
}

func (env *environment) copy() *environment {
	cpy := &environment{
		signer:    env.signer,
		state:     env.state.Copy(),
		ancestors: env.ancestors.Clone(),
		family:    env.family.Clone(),
		uncles:    env.uncles.Clone(),
		tcount:    env.tcount,
		header:    types.CopyHeader(env.header),
		receipts:  copyReceipts(env.receipts),
		sidecars:  env.sidecars, // shallow copy here intendedly
	}
	if env.gasPool != nil {
		gasPool := *env.gasPool
		cpy.gasPool = &gasPool
	}
	// The content of txs are immutable, unnecessary
	// to do the expensive deep copy for them.
	cpy.txs = make([]*types.Transaction, len(env.txs))
	copy(cpy.txs, env.txs)
	return cpy
}

// task contains all information for consensus engine sealing and result submitting.
type task struct {
	receipts  []*types.Receipt
	state     *state.StateDB
	block     *types.Block
	createdAt time.Time
	sidecars  []*types.BlobTxSidecar
}

const (
	commitInterruptNone int32 = iota
	commitInterruptNewHead
	commitInterruptResubmit
)

// newWorkReq represents a request for new sealing work submitting with relative interrupt notifier.
type newWorkReq struct {
	interrupt *int32
	noempty   bool
	timestamp int64
}

// intervalAdjust represents a resubmitting interval adjustment.
type intervalAdjust struct {
	ratio float64
	inc   bool
}

// worker is the main object which takes care of submitting new work to consensus engine
// and gathering the sealing result.
type worker struct {
	config      *Config
	chainConfig *params.ChainConfig
	engine      consensus.Engine
	eth         Backend
	chain       *core.BlockChain

	// Feeds
	pendingLogsFeed event.Feed

	// Subscriptions
	mux          *event.TypeMux
	txsCh        chan core.NewTxsEvent
	txsSub       event.Subscription
	chainHeadCh  chan core.ChainHeadEvent
	chainHeadSub event.Subscription
	chainSideCh  chan core.ChainSideEvent
	chainSideSub event.Subscription

	// Channels
	newWorkCh          chan *newWorkReq
	taskCh             chan *task
	resultCh           chan *types.Block
	startCh            chan struct{}
	exitCh             chan struct{}
	resubmitIntervalCh chan time.Duration
	resubmitAdjustCh   chan *intervalAdjust

	wg sync.WaitGroup

	current      *environment                 // An environment for current running cycle.
	localUncles  map[common.Hash]*types.Block // A set of side blocks generated locally as the possible uncle blocks.
	remoteUncles map[common.Hash]*types.Block // A set of side blocks as the possible uncle blocks.
	unconfirmed  *unconfirmedBlocks           // A set of locally mined blocks pending canonicalness confirmations.

	mu       sync.RWMutex // The lock used to protect the coinbase and extra fields
	coinbase common.Address
	extra    []byte

	pendingMu    sync.RWMutex
	pendingTasks map[common.Hash]*task

	snapshotMu       sync.RWMutex // The lock used to protect the snapshots below
	snapshotBlock    *types.Block
	snapshotReceipts types.Receipts
	snapshotState    *state.StateDB

	// atomic status counters
	running int32 // The indicator whether the consensus engine is running or not.
	newTxs  int32 // New arrival transaction count since last sealing work submitting.

	// noempty is the flag used to control whether the feature of pre-seal empty
	// block is enabled. The default value is false(pre-seal is enabled by default).
	// But in some special scenario the consensus engine will seal blocks instantaneously,
	// in this case this feature will add all empty blocks into canonical chain
	// non-stop and no real transaction will be included.
	noempty uint32

	// External functions
	isLocalBlock func(block *types.Block) bool // Function used to determine whether the specified block is mined by local miner.

	// Test hooks
	newTaskHook  func(*task)                        // Method to call upon receiving a new sealing task.
	skipSealHook func(*task) bool                   // Method to decide whether skipping the sealing.
	fullTaskHook func()                             // Method to call before pushing the full sealing task.
	resubmitHook func(time.Duration, time.Duration) // Method to call upon updating resubmitting interval.

	recentMinedBlocks *lru.Cache[uint64, []common.Hash]
}

func newWorker(config *Config, chainConfig *params.ChainConfig, engine consensus.Engine, eth Backend, mux *event.TypeMux, isLocalBlock func(*types.Block) bool, init bool) *worker {
	worker := &worker{
		config:             config,
		chainConfig:        chainConfig,
		engine:             engine,
		eth:                eth,
		mux:                mux,
		chain:              eth.BlockChain(),
		isLocalBlock:       isLocalBlock,
		localUncles:        make(map[common.Hash]*types.Block),
		remoteUncles:       make(map[common.Hash]*types.Block),
		unconfirmed:        newUnconfirmedBlocks(eth.BlockChain(), miningLogAtDepth),
		pendingTasks:       make(map[common.Hash]*task),
		txsCh:              make(chan core.NewTxsEvent, txChanSize),
		chainHeadCh:        make(chan core.ChainHeadEvent, chainHeadChanSize),
		chainSideCh:        make(chan core.ChainSideEvent, chainSideChanSize),
		newWorkCh:          make(chan *newWorkReq),
		taskCh:             make(chan *task),
		resultCh:           make(chan *types.Block, resultQueueSize),
		exitCh:             make(chan struct{}),
		startCh:            make(chan struct{}, 1),
		resubmitIntervalCh: make(chan time.Duration),
		resubmitAdjustCh:   make(chan *intervalAdjust, resubmitAdjustChanSize),
	}
	// Subscribe NewTxsEvent for tx pool
	worker.txsSub = eth.TxPool().SubscribeTransactions(worker.txsCh, false)
	// Subscribe events for blockchain
	worker.chainHeadSub = eth.BlockChain().SubscribeChainHeadEvent(worker.chainHeadCh)
	worker.chainSideSub = eth.BlockChain().SubscribeChainSideEvent(worker.chainSideCh)

	worker.recentMinedBlocks, _ = lru.New[uint64, []common.Hash](recentMinedCacheLimit)

	// Sanitize recommit interval if the user-specified one is too short.
	recommit := worker.config.Recommit
	if recommit < minRecommitInterval {
		log.Warn("Sanitizing miner recommit interval", "provided", recommit, "updated", minRecommitInterval)
		recommit = minRecommitInterval
	}

	worker.wg.Add(4)
	go worker.mainLoop()
	go worker.newWorkLoop(recommit)
	go worker.resultLoop()
	go worker.taskLoop()

	// Submit first work to initialize pending state.
	if init {
		worker.startCh <- struct{}{}
	}
	return worker
}

// setEtherbase sets the etherbase used to initialize the block coinbase field.
func (w *worker) setEtherbase(addr common.Address) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.coinbase = addr
}

func (w *worker) setGasCeil(ceil uint64) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.config.GasCeil = ceil
}

// setExtra sets the content used to initialize the block extra field.
func (w *worker) setExtra(extra []byte) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.extra = extra
}

// setRecommitInterval updates the interval for miner sealing work recommitting.
func (w *worker) setRecommitInterval(interval time.Duration) {
	w.resubmitIntervalCh <- interval
}

func (w *worker) setBlockProducerLeftover(interval time.Duration) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.config.BlockProduceLeftOver = interval
}

func (w *worker) setBlockSizeReserve(size uint64) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.config.BlockSizeReserve = size
}

// disablePreseal disables pre-sealing mining feature
func (w *worker) disablePreseal() {
	atomic.StoreUint32(&w.noempty, 1)
}

// enablePreseal enables pre-sealing mining feature
func (w *worker) enablePreseal() {
	atomic.StoreUint32(&w.noempty, 0)
}

// pending returns the pending state and corresponding block.
func (w *worker) pending() (*types.Block, *state.StateDB) {
	// return a snapshot to avoid contention on currentMu mutex
	w.snapshotMu.RLock()
	defer w.snapshotMu.RUnlock()
	if w.snapshotState == nil {
		return nil, nil
	}
	return w.snapshotBlock, w.snapshotState.Copy()
}

// pendingBlock returns pending block.
func (w *worker) pendingBlock() *types.Block {
	// return a snapshot to avoid contention on currentMu mutex
	w.snapshotMu.RLock()
	defer w.snapshotMu.RUnlock()
	return w.snapshotBlock
}

// pendingBlockAndReceipts returns pending block and corresponding receipts.
func (w *worker) pendingBlockAndReceipts() (*types.Block, types.Receipts) {
	// return a snapshot to avoid contention on currentMu mutex
	w.snapshotMu.RLock()
	defer w.snapshotMu.RUnlock()
	return w.snapshotBlock, w.snapshotReceipts
}

// start sets the running status as 1 and triggers new work submitting.
func (w *worker) start() {
	atomic.StoreInt32(&w.running, 1)
	w.startCh <- struct{}{}
}

// stop sets the running status as 0.
func (w *worker) stop() {
	atomic.StoreInt32(&w.running, 0)
}

// isRunning returns an indicator whether worker is running or not.
func (w *worker) isRunning() bool {
	return atomic.LoadInt32(&w.running) == 1
}

// close terminates all background threads maintained by the worker.
// Note the worker does not support being closed multiple times.
func (w *worker) close() {
	atomic.StoreInt32(&w.running, 0)
	close(w.exitCh)
	w.wg.Wait()
}

// recalcRecommit recalculates the resubmitting interval upon feedback.
func recalcRecommit(minRecommit, prev time.Duration, target float64, inc bool) time.Duration {
	var (
		prevF = float64(prev.Nanoseconds())
		next  float64
	)
	if inc {
		next = prevF*(1-intervalAdjustRatio) + intervalAdjustRatio*(target+intervalAdjustBias)
		max := float64(maxRecommitInterval.Nanoseconds())
		if next > max {
			next = max
		}
	} else {
		next = prevF*(1-intervalAdjustRatio) + intervalAdjustRatio*(target-intervalAdjustBias)
		min := float64(minRecommit.Nanoseconds())
		if next < min {
			next = min
		}
	}
	return time.Duration(int64(next))
}

// newWorkLoop is a standalone goroutine to submit new mining work upon received events.
func (w *worker) newWorkLoop(recommit time.Duration) {
	defer w.wg.Done()
	var (
		interrupt   *int32
		minRecommit = recommit // minimal resubmit interval specified by user.
		timestamp   int64      // timestamp for each round of mining.
	)

	timer := time.NewTimer(0)
	defer timer.Stop()
	<-timer.C // discard the initial tick

	// commit aborts in-flight transaction execution with given signal and resubmits a new one.
	commit := func(noempty bool, s int32) {
		if interrupt != nil {
			atomic.StoreInt32(interrupt, s)
		}
		interrupt = new(int32)
		select {
		case w.newWorkCh <- &newWorkReq{interrupt: interrupt, noempty: noempty, timestamp: timestamp}:
		case <-w.exitCh:
			return
		}
		timer.Reset(recommit)
		atomic.StoreInt32(&w.newTxs, 0)
	}
	// clearPending cleans the stale pending tasks.
	clearPending := func(number uint64) {
		w.pendingMu.Lock()
		for h, t := range w.pendingTasks {
			if t.block.NumberU64()+staleThreshold <= number {
				delete(w.pendingTasks, h)
			}
		}
		w.pendingMu.Unlock()
	}

	for {
		select {
		case <-w.startCh:
			clearPending(w.chain.CurrentBlock().NumberU64())
			timestamp = time.Now().Unix()
			commit(false, commitInterruptNewHead)

		case head := <-w.chainHeadCh:
			clearPending(head.Block.NumberU64())
			timestamp = time.Now().Unix()
			commit(false, commitInterruptNewHead)

		case <-timer.C:
			if w.chainConfig.Consortium != nil {
				continue
			}

			// If mining is running resubmit a new work cycle periodically to pull in
			// higher priced transactions. Disable this overhead for pending blocks.
			if w.isRunning() && (w.chainConfig.Clique == nil || w.chainConfig.Clique.Period > 0) {
				// Short circuit if no new transaction arrives.
				if atomic.LoadInt32(&w.newTxs) == 0 {
					timer.Reset(recommit)
					continue
				}
				commit(true, commitInterruptResubmit)
			}

		case interval := <-w.resubmitIntervalCh:
			// Adjust resubmit interval explicitly by user.
			if interval < minRecommitInterval {
				log.Warn("Sanitizing miner recommit interval", "provided", interval, "updated", minRecommitInterval)
				interval = minRecommitInterval
			}
			log.Info("Miner recommit interval update", "from", minRecommit, "to", interval)
			minRecommit, recommit = interval, interval

			if w.resubmitHook != nil {
				w.resubmitHook(minRecommit, recommit)
			}

		case adjust := <-w.resubmitAdjustCh:
			// Adjust resubmit interval by feedback.
			if adjust.inc {
				before := recommit
				target := float64(recommit.Nanoseconds()) / adjust.ratio
				recommit = recalcRecommit(minRecommit, recommit, target, true)
				log.Trace("Increase miner recommit interval", "from", before, "to", recommit)
			} else {
				before := recommit
				recommit = recalcRecommit(minRecommit, recommit, float64(minRecommit.Nanoseconds()), false)
				log.Trace("Decrease miner recommit interval", "from", before, "to", recommit)
			}

			if w.resubmitHook != nil {
				w.resubmitHook(minRecommit, recommit)
			}

		case <-w.exitCh:
			return
		}
	}
}

// mainLoop is a standalone goroutine to regenerate the sealing task based on the received event.
func (w *worker) mainLoop() {
	defer w.wg.Done()
	defer w.txsSub.Unsubscribe()
	defer w.chainHeadSub.Unsubscribe()
	defer w.chainSideSub.Unsubscribe()
	defer func() {
		if w.current != nil && w.current.state != nil {
			w.current.state.StopPrefetcher()
		}
	}()

	for {
		select {
		case req := <-w.newWorkCh:
			w.commitNewWork(req.interrupt, req.noempty, req.timestamp)

		case ev := <-w.chainSideCh:
			// Don't handle uncle logic in Consortium
			if w.chainConfig.Consortium != nil {
				continue
			}
			// Short circuit for duplicate side blocks
			if _, exist := w.localUncles[ev.Block.Hash()]; exist {
				continue
			}
			if _, exist := w.remoteUncles[ev.Block.Hash()]; exist {
				continue
			}
			// Add side block to possible uncle block set depending on the author.
			if w.isLocalBlock != nil && w.isLocalBlock(ev.Block) {
				w.localUncles[ev.Block.Hash()] = ev.Block
			} else {
				w.remoteUncles[ev.Block.Hash()] = ev.Block
			}
			// If our mining block contains less than 2 uncle blocks,
			// add the new uncle block if valid and regenerate a mining block.
			if w.isRunning() && w.current != nil && w.current.uncles.Cardinality() < 2 {
				start := time.Now()
				if err := w.commitUncle(w.current, ev.Block.Header()); err == nil {
					var uncles []*types.Header
					w.current.uncles.Each(func(item interface{}) bool {
						hash, ok := item.(common.Hash)
						if !ok {
							return false
						}
						uncle, exist := w.localUncles[hash]
						if !exist {
							uncle, exist = w.remoteUncles[hash]
						}
						if !exist {
							return false
						}
						uncles = append(uncles, uncle.Header())
						return false
					})
					w.commit(uncles, nil, true, start)
				}
			}

		case ev := <-w.txsCh:
			// Apply transactions to the pending state if we're not mining.
			//
			// Note all transactions received may not be continuous with transactions
			// already included in the current mining block. These transactions will
			// be automatically eliminated.
			if !w.isRunning() && w.current != nil {
				// If block is already full, abort
				if gp := w.current.gasPool; gp != nil && gp.Gas() < params.TxGas {
					continue
				}
				w.mu.RLock()
				coinbase := w.coinbase
				w.mu.RUnlock()

				txs := make(map[common.Address][]*txpool.LazyTransaction)
				for _, tx := range ev.Txs {
					acc, _ := types.Sender(w.current.signer, tx)
					txs[acc] = append(txs[acc], &txpool.LazyTransaction{
						Tx:        tx,
						Hash:      tx.Hash(),
						Time:      tx.Time(),
						GasFeeCap: uint256.MustFromBig(tx.GasFeeCap()),
						GasTipCap: uint256.MustFromBig(tx.GasTipCap()),
						Gas:       tx.Gas(),
						BlobGas:   tx.BlobGas(),
					})
				}
				txset := NewTransactionsByPriceAndNonce(w.current.signer, txs, w.current.header.BaseFee)
				emptyset := NewTransactionsByPriceAndNonce(w.current.signer, make(map[common.Address][]*txpool.LazyTransaction), w.current.header.BaseFee)
				tcount := w.current.tcount
				w.commitTransactions(txset, emptyset, coinbase, nil)
				// Only update the snapshot if any new transactons were added
				// to the pending block
				if tcount != w.current.tcount {
					w.updateSnapshot()
				}
			} else {
				// Special case, if the consensus engine is 0 period clique(dev mode),
				// submit mining work here since all empty submission will be rejected
				// by clique. Of course the advance sealing(empty submission) is disabled.
				if w.chainConfig.Clique != nil && w.chainConfig.Clique.Period == 0 {
					w.commitNewWork(nil, true, time.Now().Unix())
				}
			}
			atomic.AddInt32(&w.newTxs, int32(len(ev.Txs)))

		// System stopped
		case <-w.exitCh:
			return
		case <-w.txsSub.Err():
			return
		case <-w.chainHeadSub.Err():
			return
		case <-w.chainSideSub.Err():
			return
		}
	}
}

// taskLoop is a standalone goroutine to fetch sealing task from the generator and
// push them to consensus engine.
func (w *worker) taskLoop() {
	defer w.wg.Done()
	var (
		stopCh chan struct{}
		prev   common.Hash
	)

	// interrupt aborts the in-flight sealing task.
	interrupt := func() {
		if stopCh != nil {
			close(stopCh)
			stopCh = nil
		}
	}
	for {
		select {
		case task := <-w.taskCh:
			if w.newTaskHook != nil {
				w.newTaskHook(task)
			}
			// Reject duplicate sealing work due to resubmitting.
			sealHash := w.engine.SealHash(task.block.Header())
			if sealHash == prev {
				continue
			}
			// Interrupt previous sealing operation
			interrupt()
			stopCh, prev = make(chan struct{}), sealHash

			if w.skipSealHook != nil && w.skipSealHook(task) {
				continue
			}
			w.pendingMu.Lock()
			w.pendingTasks[sealHash] = task
			w.pendingMu.Unlock()

			if err := w.engine.Seal(w.chain, task.block, w.resultCh, stopCh); err != nil {
				log.Warn("Block sealing failed", "err", err, "height", task.block.NumberU64())
				w.pendingMu.Lock()
				delete(w.pendingTasks, sealHash)
				w.pendingMu.Unlock()
			}
		case <-w.exitCh:
			interrupt()
			return
		}
	}
}

// resultLoop is a standalone goroutine to handle sealing result submitting
// and flush relative data to the database.
func (w *worker) resultLoop() {
	defer w.wg.Done()
	for {
		select {
		case block := <-w.resultCh:
			// Short circuit when receiving empty result.
			if block == nil {
				continue
			}
			// Short circuit when receiving duplicate result caused by resubmitting.
			if w.chain.HasBlock(block.Hash(), block.NumberU64()) {
				continue
			}
			var (
				sealhash = w.engine.SealHash(block.Header())
				hash     = block.Hash()
			)
			w.pendingMu.RLock()
			task, exist := w.pendingTasks[sealhash]
			w.pendingMu.RUnlock()
			if !exist {
				log.Error("Block found but no relative pending task", "number", block.Number(), "sealhash", sealhash, "hash", hash)
				continue
			}

			if w.chainConfig.IsConsortiumV2(block.Number()) {
				if parentsHash, ok := w.recentMinedBlocks.Get(block.NumberU64()); ok {
					isDoubleSign := false
					for _, parent := range parentsHash {
						if block.ParentHash() == parent {
							isDoubleSign = true
							log.Info("Possibly double sign, drop block", "block", block.Number(), "hash", block.Hash(), "parent", block.ParentHash())
						}
					}
					// Don't broadcast and write this block to chain,
					// continue receiving next sealed block from result channel
					if isDoubleSign {
						continue
					}
					parentsHash = append(parentsHash, block.ParentHash())
					w.recentMinedBlocks.Add(block.NumberU64(), parentsHash)
				} else {
					w.recentMinedBlocks.Add(block.NumberU64(), []common.Hash{block.ParentHash()})
				}
			}

			// Different block could share same sealhash, deep copy here to prevent write-write conflict.
			var (
				receipts = make([]*types.Receipt, len(task.receipts))
				logs     []*types.Log
			)
			for i, taskReceipt := range task.receipts {
				receipt := new(types.Receipt)
				receipts[i] = receipt
				*receipt = *taskReceipt

				// add block location fields
				receipt.BlockHash = hash
				receipt.BlockNumber = block.Number()
				receipt.TransactionIndex = uint(i)

				// Update the block hash in all logs since it is now available and not when the
				// receipt/log of individual transactions were created.
				receipt.Logs = make([]*types.Log, len(taskReceipt.Logs))
				for i, taskLog := range taskReceipt.Logs {
					log := new(types.Log)
					receipt.Logs[i] = log
					*log = *taskLog
					log.BlockHash = hash
				}
				logs = append(logs, receipt.Logs...)
			}
			// Commit block and state to database.
			_, err := w.chain.WriteBlockWithState(block, receipts, logs, task.state, true, task.sidecars)
			if err != nil {
				log.Error("Failed writing block to chain", "err", err)
				continue
			}
			log.Info("Successfully sealed new block", "number", block.Number(), "sealhash", sealhash, "hash", hash,
				"elapsed", common.PrettyDuration(time.Since(task.createdAt)), "txs", len(block.Transactions()))

			// Broadcast the block and announce chain insertion event
			w.mux.Post(core.NewMinedBlockEvent{Block: block, Sidecars: task.sidecars})

			// Insert the block into the set of pending ones to resultLoop for confirmations
			w.unconfirmed.Insert(block.NumberU64(), block.Hash())

		case <-w.exitCh:
			return
		}
	}
}

// makeCurrent creates a new environment for the current cycle.
func (w *worker) makeCurrent(parent *types.Block, header *types.Header) error {
	// Retrieve the parent state to execute on top and start a prefetcher for
	// the miner to speed block sealing up a bit
	state, err := w.chain.StateAt(parent.Root())
	if err != nil {
		return err
	}
	state.StartPrefetcher("miner")

	env := &environment{
		signer:             types.MakeSigner(w.chainConfig, header.Number),
		state:              state,
		ancestors:          mapset.NewSet(),
		family:             mapset.NewSet(),
		uncles:             mapset.NewSet(),
		header:             header,
		estimatedBlockSize: 0,
	}
	// when 08 is processed ancestors contain 07 (quick block)
	for _, ancestor := range w.chain.GetBlocksFromHash(parent.Hash(), 7) {
		for _, uncle := range ancestor.Uncles() {
			env.family.Add(uncle.Hash())
		}
		env.family.Add(ancestor.Hash())
		env.ancestors.Add(ancestor.Hash())
	}
	// Keep track of transactions which return errors so they can be removed
	env.tcount = 0

	// Swap out the old work with the new one, terminating any leftover prefetcher
	// processes in the mean time and starting a new one.
	if w.current != nil && w.current.state != nil {
		w.current.state.StopPrefetcher()
	}
	w.current = env
	return nil
}

// commitUncle adds the given block to uncle block set, returns error if failed to add.
func (w *worker) commitUncle(env *environment, uncle *types.Header) error {
	hash := uncle.Hash()
	if env.uncles.Contains(hash) {
		return errors.New("uncle not unique")
	}
	if env.header.ParentHash == uncle.ParentHash {
		return errors.New("uncle is sibling")
	}
	if !env.ancestors.Contains(uncle.ParentHash) {
		return errors.New("uncle's parent unknown")
	}
	if env.family.Contains(hash) {
		return errors.New("uncle already included")
	}
	env.uncles.Add(uncle.Hash())
	return nil
}

// updateSnapshot updates pending snapshot block and state.
// Note this function assumes the current variable is thread safe.
func (w *worker) updateSnapshot() {
	w.snapshotMu.Lock()
	defer w.snapshotMu.Unlock()

	var uncles []*types.Header
	w.current.uncles.Each(func(item interface{}) bool {
		hash, ok := item.(common.Hash)
		if !ok {
			return false
		}
		uncle, exist := w.localUncles[hash]
		if !exist {
			uncle, exist = w.remoteUncles[hash]
		}
		if !exist {
			return false
		}
		uncles = append(uncles, uncle.Header())
		return false
	})

	w.snapshotBlock = types.NewBlock(
		w.current.header,
		w.current.txs,
		uncles,
		w.current.receipts,
		trie.NewStackTrie(nil),
	)
	w.snapshotReceipts = copyReceipts(w.current.receipts)
	w.snapshotState = w.current.state.Copy()
}

func (w *worker) commitTransaction(tx *types.Transaction, coinbase common.Address, receiptProcessor core.ReceiptProcessor) ([]*types.Log, error) {
	txType := tx.Type()
	if txType == types.BlobTxType {
		if tx.BlobTxSidecar() == nil {
			return nil, errors.New("blob transaction with empty sidecar in miner")
		}
		if w.current.header.BlobGasUsed == nil {
			return nil, errors.New("blob transaction when not Cancun")
		}
	}

	snap := w.current.state.Snapshot()

	receipt, _, err := core.ApplyTransaction(
		w.chainConfig,
		w.chain,
		&coinbase,
		w.current.gasPool,
		w.current.state,
		w.current.header,
		tx,
		&w.current.header.GasUsed,
		*w.chain.GetVMConfig(),
		receiptProcessor,
	)
	if err != nil {
		w.current.state.RevertToSnapshot(snap)
		return nil, err
	}
	w.current.txs = append(w.current.txs, tx.WithoutBlobTxSidecar())
	w.current.receipts = append(w.current.receipts, receipt)
	w.current.tcount++
	w.current.estimatedBlockSize += uint64(tx.Size())
	if tx.Type() == types.BlobTxType {
		*w.current.header.BlobGasUsed += tx.BlobGas()
		w.current.sidecars = append(w.current.sidecars, tx.BlobTxSidecar())
	}

	return receipt.Logs, nil
}

func (w *worker) commitTransactions(plainTxs, blobTxs *TransactionsByPriceAndNonce, coinbase common.Address, interrupt *int32) bool {
	// Short circuit if current is nil
	if w.current == nil {
		return true
	}

	var reservedGas uint64 = 0
	if w.chainConfig.Consortium != nil {
		if w.current.header.Number.Uint64()%w.chainConfig.Consortium.EpochV2 == w.chainConfig.Consortium.EpochV2-1 {
			reservedGas = params.ReservedGasForCheckpointSystemTransactions
		} else {
			reservedGas = params.ReservedGasForNormalSystemTransactions
		}
	}
	gasLimit := w.current.header.GasLimit
	if w.current.gasPool == nil {
		w.current.gasPool = new(core.GasPool).AddGas(gasLimit)

		// If the gas pool is newly created, reserve some gas for system transactions
		if w.chainConfig.Consortium != nil {
			if err := w.current.gasPool.SubGas(reservedGas); err != nil {
				log.Error(
					"Failed to reserve gas for system transactions",
					"pool", w.current.gasPool,
					"reserve", reservedGas,
					"error", err,
				)
				return true
			}
		}
	}

	var coalescedLogs []*types.Log
	var timer *time.Timer

	isCancun := w.chainConfig.IsCancun(w.current.header.Number)
	if !isCancun && blobTxs.Size() != 0 {
		log.Error("Get blob transaction when not Cancun", "block", w.current.header.Number.Uint64())
		// Clear all blob transactions and continue committing plain transactions
		blobTxs.Clear()
	}

	// This timer is only shipped after Buba hardfork
	// When it is nearly the time an empty block is produced,
	// we break the commit transactions process, allow the
	// block with transactions to be produced and abort the
	// empty block producer.
	if w.chainConfig.IsBuba(w.current.header.Number) {
		duration := time.Until(time.Unix(int64(w.current.header.Time), 0)) - w.config.BlockProduceLeftOver
		timer = time.NewTimer(duration)
	}

	bloomProcessor := core.NewAsyncReceiptBloomGenerator(plainTxs.Size() + blobTxs.Size())
	defer bloomProcessor.Close()

Loop:
	for {
		if timer != nil {
			select {
			case <-timer.C:
				break Loop
			default:
			}
		}

		// In the following three cases, we will interrupt the execution of the transaction.
		// (1) new head block event arrival, the interrupt signal is 1
		// (2) worker start or restart, the interrupt signal is 1
		// (3) worker recreate the mining block with any newly arrived transactions, the interrupt signal is 2.
		// For the first two cases, the semi-finished work will be discarded.
		// For the third case, the semi-finished work will be submitted to the consensus engine.
		if interrupt != nil && atomic.LoadInt32(interrupt) != commitInterruptNone {
			// Notify resubmit loop to increase resubmitting interval due to too frequent commits.
			if atomic.LoadInt32(interrupt) == commitInterruptResubmit {
				ratio := float64(gasLimit-w.current.gasPool.Gas()) / float64(gasLimit)
				if ratio < 0.1 {
					ratio = 0.1
				}
				w.resubmitAdjustCh <- &intervalAdjust{
					ratio: ratio,
					inc:   true,
				}
			}
			return atomic.LoadInt32(interrupt) == commitInterruptNewHead
		}
		// If we don't have enough gas for any further transactions then we're done
		if w.current.gasPool.Gas() < params.TxGas {
			log.Trace("Not enough gas for further transactions", "have", w.current.gasPool, "want", params.TxGas)
			break
		}

		if isCancun && blobTxs.Size() != 0 && params.MaxBlobGasPerBlock <= *w.current.header.BlobGasUsed {
			log.Debug("Max blob reached, stop committing blob transactions")
			blobTxs.Clear()
		}

		if w.current.estimatedBlockSize+w.config.BlockSizeReserve > maxBlockSize {
			log.Debug("Estimated block size is too big", "estimated size", w.current.estimatedBlockSize)
			break
		}

		// Retrieve the next transaction and abort if all done
		var (
			selectedTx  *txpool.LazyTransaction
			selectedTxs *TransactionsByPriceAndNonce
		)
		plainTx, plainTip := plainTxs.Peek()
		blobTx, blobTip := blobTxs.Peek()
		if plainTx == nil {
			selectedTx, selectedTxs = blobTx, blobTxs
		} else if blobTx == nil {
			selectedTx, selectedTxs = plainTx, plainTxs
		} else {
			if plainTip.Cmp(blobTip) < 0 {
				selectedTx, selectedTxs = blobTx, blobTxs
			} else {
				selectedTx, selectedTxs = plainTx, plainTxs
			}
		}

		if selectedTx == nil {
			break
		}
		txs := selectedTxs
		// If we don't have enough space for the next transaction, skip the account.
		if w.current.gasPool.Gas() < selectedTx.Gas {
			log.Trace("Not enough gas left for transaction", "hash", selectedTx.Hash, "left", w.current.gasPool.Gas(), "needed", selectedTx.Gas)
			txs.Pop()
			continue
		}

		if isCancun {
			if left := params.MaxBlobGasPerBlock - *w.current.header.BlobGasUsed; left < selectedTx.BlobGas {
				log.Trace("Not enough blob gas left for transaction", "hash", selectedTx.Hash, "left", left, "needed", selectedTx.BlobGas)
				txs.Pop()
				continue
			}
		}

		tx := selectedTx.Resolve()
		if tx == nil {
			log.Trace("Transaction is evicted from pool", "hash", selectedTx.Hash)
			txs.Pop()
			continue
		}

		// Error may be ignored here. The error has already been checked
		// during transaction acceptance is the transaction pool.
		//
		// We use the eip155 signer regardless of the current hf.
		from, _ := types.Sender(w.current.signer, tx)
		// Check whether the tx is replay protected. If we're not in the EIP155 hf
		// phase, start ignoring the sender until we do.
		if tx.Protected() && !w.chainConfig.IsEIP155(w.current.header.Number) {
			log.Trace("Ignoring reply protected transaction", "hash", tx.Hash(), "eip155", w.chainConfig.EIP155Block)

			txs.Pop()
			continue
		}
		// Start executing the transaction
		w.current.state.SetTxContext(tx.Hash(), w.current.tcount)

		logs, err := w.commitTransaction(tx, coinbase, bloomProcessor)
		switch {
		case errors.Is(err, core.ErrGasLimitReached):
			// Pop the current out-of-gas transaction without shifting in the next from the account
			log.Trace("Gas limit exceeded for current block", "sender", from)
			txs.Pop()

		case errors.Is(err, core.ErrNonceTooLow):
			// New head notification data race between the transaction pool and miner, shift
			log.Trace("Skipping transaction with low nonce", "sender", from, "nonce", tx.Nonce())
			txs.Shift()

		case errors.Is(err, core.ErrNonceTooHigh):
			// Reorg notification data race between the transaction pool and miner, skip account =
			log.Trace("Skipping account with hight nonce", "sender", from, "nonce", tx.Nonce())
			txs.Pop()

		case errors.Is(err, nil):
			// Everything ok, collect the logs and shift in the next transaction from the same account
			coalescedLogs = append(coalescedLogs, logs...)
			txs.Shift()

		case errors.Is(err, types.ErrTxTypeNotSupported):
			// Pop the unsupported transaction without shifting in the next from the account
			log.Trace("Skipping unsupported transaction type", "sender", from, "type", tx.Type())
			txs.Pop()

		default:
			// Strange error, discard the transaction and get the next in line (note, the
			// nonce-too-high clause will prevent us from executing in vain).
			log.Debug("Transaction failed, account skipped", "hash", tx.Hash(), "err", err)
			txs.Shift()
		}
	}

	if !w.isRunning() && len(coalescedLogs) > 0 {
		// We don't push the pendingLogsEvent while we are mining. The reason is that
		// when we are mining, the worker will regenerate a mining block every 3 seconds.
		// In order to avoid pushing the repeated pendingLog, we disable the pending log pushing.

		// make a copy, the state caches the logs and these logs get "upgraded" from pending to mined
		// logs by filling in the block hash when the block was mined by the local miner. This can
		// cause a race condition if a log was "upgraded" before the PendingLogsEvent is processed.
		cpy := make([]*types.Log, len(coalescedLogs))
		for i, l := range coalescedLogs {
			cpy[i] = new(types.Log)
			*cpy[i] = *l
		}
		w.pendingLogsFeed.Send(cpy)
	}
	// Notify resubmit loop to decrease resubmitting interval if current interval is larger
	// than the user-specified one.
	if interrupt != nil {
		w.resubmitAdjustCh <- &intervalAdjust{inc: false}
	}
	return false
}

// commitNewWork generates several new sealing tasks based on the parent block.
func (w *worker) commitNewWork(interrupt *int32, noempty bool, timestamp int64) {
	w.mu.RLock()
	defer w.mu.RUnlock()

	tstart := time.Now()

	var (
		parent     *types.Block
		fastCommit bool
	)

	parent = w.chain.CurrentBlock()
	if consortiumEngine, ok := w.engine.(*consortium.Consortium); ok {
		if w.chainConfig.IsOlek(parent.Number()) {
			parent, fastCommit = consortiumEngine.GetBestParentBlock(w.chain)
		}
	}

	if parent.Time() >= uint64(timestamp) {
		timestamp = int64(parent.Time() + 1)
	}
	num := parent.Number()
	header := &types.Header{
		ParentHash: parent.Hash(),
		Number:     num.Add(num, common.Big1),
		GasLimit:   core.CalcGasLimit(parent.GasLimit(), w.config.GasCeil),
		Extra:      w.extra,
		Time:       uint64(timestamp),
	}
	// Set baseFee if we are on an EIP-1559 chain
	if w.chainConfig.IsLondon(header.Number) {
		header.BaseFee = eip1559.CalcBaseFee(w.chainConfig, parent.Header())
	}
	if w.chainConfig.IsCancun(header.Number) {
		var excessBlobGas uint64
		if w.chainConfig.IsCancun(parent.Number()) {
			excessBlobGas = eip4844.CalcExcessBlobGas(*parent.ExcessBlobGas(), *parent.BlobGasUsed())
		} else {
			excessBlobGas = eip4844.CalcExcessBlobGas(0, 0)
		}
		header.BlobGasUsed = new(uint64)
		header.ExcessBlobGas = &excessBlobGas
	}
	// Only set the coinbase if our consensus engine is running (avoid spurious block rewards)
	if w.isRunning() {
		if w.coinbase == (common.Address{}) {
			log.Error("Refusing to mine without etherbase")
			return
		}
		header.Coinbase = w.coinbase
	}
	if err := w.engine.Prepare(w.chain, header); err != nil {
		log.Error("Failed to prepare header for mining", "err", err, "number", header.Number.Uint64())
		return
	}
	// If we are care about TheDAO hard-fork check whether to override the extra-data or not
	if daoBlock := w.chainConfig.DAOForkBlock; daoBlock != nil {
		// Check whether the block is among the fork extra-override range
		limit := new(big.Int).Add(daoBlock, params.DAOForkExtraRange)
		if header.Number.Cmp(daoBlock) >= 0 && header.Number.Cmp(limit) < 0 {
			// Depending whether we support or oppose the fork, override differently
			if w.chainConfig.DAOForkSupport {
				header.Extra = common.CopyBytes(params.DAOForkBlockExtra)
			} else if bytes.Equal(header.Extra, params.DAOForkBlockExtra) {
				header.Extra = []byte{} // If miner opposes, don't let it use the reserved extra-data
			}
		}
	}
	// Could potentially happen if starting to mine in an odd state.
	err := w.makeCurrent(parent, header)
	if err != nil {
		log.Error("Failed to create mining context", "err", err)
		return
	}
	// Create the current work task and check any fork transitions needed
	env := w.current
	if w.chainConfig.DAOForkSupport && w.chainConfig.DAOForkBlock != nil && w.chainConfig.DAOForkBlock.Cmp(header.Number) == 0 {
		misc.ApplyDAOHardFork(env.state)
	}

	// Accumulate the uncles for the current block
	uncles := make([]*types.Header, 0, 2)
	commitUncles := func(blocks map[common.Hash]*types.Block) {
		// Clean up stale uncle blocks first
		for hash, uncle := range blocks {
			if uncle.NumberU64()+staleThreshold <= header.Number.Uint64() {
				delete(blocks, hash)
			}
		}
		for hash, uncle := range blocks {
			if len(uncles) == 2 {
				break
			}
			if err := w.commitUncle(env, uncle.Header()); err != nil {
				log.Trace("Possible uncle rejected", "hash", hash, "reason", err)
			} else {
				log.Debug("Committing new uncle to block", "hash", hash)
				uncles = append(uncles, uncle.Header())
			}
		}
	}
	// Prefer to locally generated uncle
	commitUncles(w.localUncles)
	commitUncles(w.remoteUncles)

	// Fast commit an empty block which has block number smaller
	// than current block in canonical chain but have greater
	// difficulty
	if fastCommit {
		w.commit(uncles, nil, true, tstart)
		return
	}

	// Create an empty block based on temporary copied state for
	// sealing in advance without waiting block execution finished.
	if !noempty && atomic.LoadUint32(&w.noempty) == 0 {
		w.commit(uncles, nil, false, tstart)
	}

	filter := txpool.PendingFilter{
		EnforceTip: true,
	}
	if w.current.header.BaseFee != nil {
		filter.BaseFee = uint256.MustFromBig(w.current.header.BaseFee)
	}
	if w.current.header.ExcessBlobGas != nil {
		filter.BlobFee = uint256.MustFromBig(eip4844.CalcBlobFee(*w.current.header.ExcessBlobGas))
	}
	filter.OnlyPlainTxs, filter.OnlyBlobTxs = true, false
	pendingPlainTxs := w.eth.TxPool().Pending(&filter)

	filter.OnlyPlainTxs, filter.OnlyBlobTxs = false, true
	pendingBlobTxs := w.eth.TxPool().Pending(&filter)

	// Short circuit if there is no available pending transactions.
	// But if we disable empty precommit already, ignore it. Since
	// empty block is necessary to keep the liveness of the network.
	if len(pendingPlainTxs) == 0 && len(pendingBlobTxs) == 0 && atomic.LoadUint32(&w.noempty) == 0 {
		w.updateSnapshot()
		return
	}

	// Split the pending transactions into locals and remotes
	localAccounts := w.eth.TxPool().Locals()
	localPlainTxs, remotePlainTxs := make(map[common.Address][]*txpool.LazyTransaction), pendingPlainTxs
	for _, account := range localAccounts {
		if txs := remotePlainTxs[account]; len(txs) > 0 {
			delete(remotePlainTxs, account)
			localPlainTxs[account] = txs
		}
	}
	localBlobTxs, remoteBlobTxs := make(map[common.Address][]*txpool.LazyTransaction), pendingBlobTxs
	for _, account := range localAccounts {
		if txs := remoteBlobTxs[account]; len(txs) > 0 {
			delete(remoteBlobTxs, account)
			localBlobTxs[account] = txs
		}
	}
	if len(localPlainTxs) > 0 || len(localBlobTxs) > 0 {
		plainTxs := NewTransactionsByPriceAndNonce(w.current.signer, localPlainTxs, header.BaseFee)
		blobTxs := NewTransactionsByPriceAndNonce(w.current.signer, localBlobTxs, header.BaseFee)
		if w.commitTransactions(plainTxs, blobTxs, w.coinbase, interrupt) {
			return
		}
	}
	if len(remotePlainTxs) > 0 || len(remoteBlobTxs) > 0 {
		plainTxs := NewTransactionsByPriceAndNonce(w.current.signer, remotePlainTxs, header.BaseFee)
		blobTxs := NewTransactionsByPriceAndNonce(w.current.signer, remoteBlobTxs, header.BaseFee)
		if w.commitTransactions(plainTxs, blobTxs, w.coinbase, interrupt) {
			return
		}
	}
	w.commit(uncles, w.fullTaskHook, true, tstart)
}

// commit runs any post-transaction state modifications, assembles the final block
// and commits new work if consensus engine is running.
func (w *worker) commit(uncles []*types.Header, interval func(), update bool, start time.Time) error {
	if w.isRunning() {
		// Deep copy receipts here to avoid interaction between different tasks.
		env := w.current.copy()
		// As consortium does not use uncles, we don't care about copying uncles here
		block, receipts, err := w.engine.FinalizeAndAssemble(w.chain, env.header, env.state, env.txs, uncles, env.receipts)
		if err != nil {
			log.Error("Failed to FinalizeAndAssemble a block", "error", err)
			return err
		}

		if interval != nil {
			interval()
		}
		select {
		case w.taskCh <- &task{receipts: receipts, state: env.state, block: block, createdAt: time.Now(), sidecars: env.sidecars}:
			w.unconfirmed.Shift(block.NumberU64() - 1)
			log.Info("Commit new mining work", "number", block.Number(), "sealhash", w.engine.SealHash(block.Header()),
				"uncles", len(uncles), "txs", w.current.tcount,
				"gas", block.GasUsed(), "fees", totalFees(block, receipts),
				"elapsed", common.PrettyDuration(time.Since(start)))

		case <-w.exitCh:
			log.Info("Worker has exited")
		}
	}
	if update {
		w.updateSnapshot()
	}
	return nil
}

// copyReceipts makes a deep copy of the given receipts.
func copyReceipts(receipts []*types.Receipt) []*types.Receipt {
	result := make([]*types.Receipt, len(receipts))
	for i, l := range receipts {
		cpy := *l
		result[i] = &cpy
	}
	return result
}

// postSideBlock fires a side chain event, only use it for testing.
func (w *worker) postSideBlock(event core.ChainSideEvent) {
	select {
	case w.chainSideCh <- event:
	case <-w.exitCh:
	}
}

// totalFees computes total consumed miner fees in ETH. Block transactions and receipts have to have the same order.
func totalFees(block *types.Block, receipts []*types.Receipt) *big.Float {
	feesWei := new(big.Int)
	for i, tx := range block.Transactions() {
		minerFee, _ := tx.EffectiveGasTip(block.BaseFee())
		feesWei.Add(feesWei, new(big.Int).Mul(new(big.Int).SetUint64(receipts[i].GasUsed), minerFee))
	}
	return new(big.Float).Quo(new(big.Float).SetInt(feesWei), new(big.Float).SetInt(big.NewInt(params.Ether)))
}
