package tracer

import (
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
)

var (
	callStackRecorderInit sync.Once
	recorder              *callStackRecorder
)

func WithTracer(hooks *tracing.Hooks) *tracing.Hooks {
	callStackRecorderInit.Do(func() {
		recorder = newRecorder()
	})
	return recorder.withHooks(hooks)
}

func GetParentOrder() uint64 {
	if recorder == nil {
		return 0
	}
	return recorder.getParentOrder()
}

type callStackRecorder struct {
	hooks  *tracing.Hooks
	orders []uint64
}

func newRecorder() *callStackRecorder {
	return &callStackRecorder{orders: make([]uint64, 1)}
}

func (c *callStackRecorder) withHooks(hooks *tracing.Hooks) *tracing.Hooks {
	c.hooks = hooks
	if hooks == nil {
		return &tracing.Hooks{
			OnEnter: c.onEnter,
			OnExit:  c.onExit,
		}
	}
	return &tracing.Hooks{
		OnTxStart:           hooks.OnTxStart,
		OnTxEnd:             hooks.OnTxEnd,
		OnEnter:             c.onEnter,
		OnExit:              c.onExit,
		OnOpcode:            hooks.OnOpcode,
		OnFault:             hooks.OnFault,
		OnGasChange:         hooks.OnGasChange,
		OnBlockchainInit:    hooks.OnBlockchainInit,
		OnClose:             hooks.OnClose,
		OnBlockStart:        hooks.OnBlockStart,
		OnBlockEnd:          hooks.OnBlockEnd,
		OnSkippedBlock:      hooks.OnSkippedBlock,
		OnGenesisBlock:      hooks.OnGenesisBlock,
		OnSystemCallStart:   hooks.OnSystemCallEnd,
		OnSystemCallStartV2: hooks.OnSystemCallStartV2,
		OnSystemCallEnd:     hooks.OnSystemCallEnd,
		OnBalanceChange:     hooks.OnBalanceChange,
		OnNonceChange:       hooks.OnNonceChange,
		OnNonceChangeV2:     hooks.OnNonceChangeV2,
		OnCodeChange:        hooks.OnCodeChange,
		OnStorageChange:     hooks.OnStorageChange,
		OnLog:               hooks.OnLog,
		OnBlockHashRead:     hooks.OnBlockHashRead,
		GetParentOrder:      nil,
	}
}

func (c *callStackRecorder) onEnter(depth int, typ byte, from common.Address, to common.Address, input []byte, gas uint64, value *big.Int, order uint64) {
	c.orders = append(c.orders, order)
	// call underlying hooks if any
	if c.hooks != nil {
		c.hooks.OnEnter(depth, typ, from, to, input, gas, value, order)
	}
}

func (c *callStackRecorder) onExit(depth int, output []byte, gasUsed uint64, err error, reverted bool) {
	size := len(c.orders)
	if size <= 1 {
		return
	}
	// pop call
	c.orders = c.orders[:size-1]
	if c.hooks != nil {
		c.hooks.OnExit(depth, output, gasUsed, err, reverted)
	}
}

func (c *callStackRecorder) getParentOrder() uint64 {
	if len(c.orders) <= 1 {
		return 0
	}
	return c.orders[len(c.orders)-2]
}
