package tracer

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
)

type CallStackRecorder struct {
	hooks  *tracing.Hooks
	orders []uint64
}

func NewCallTracerOrder(hooks *tracing.Hooks) (*CallStackRecorder, *tracing.Hooks) {
	t := &CallStackRecorder{orders: make([]uint64, 1), hooks: hooks}
	if hooks == nil {
		return t, &tracing.Hooks{
			OnEnter: t.onEnter,
			OnExit:  t.onExit,
		}
	}
	return t, &tracing.Hooks{
		OnTxStart:           hooks.OnTxStart,
		OnTxEnd:             hooks.OnTxEnd,
		OnEnter:             t.onEnter,
		OnExit:              t.onExit,
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

func (t *CallStackRecorder) onEnter(depth int, typ byte, from common.Address, to common.Address, input []byte, gas uint64, value *big.Int, order uint64) {
	t.orders = append(t.orders, order)
	// call underlying hooks if any
	if t.hooks != nil {
		t.hooks.OnEnter(depth, typ, from, to, input, gas, value, order)
	}
}

func (t *CallStackRecorder) onExit(depth int, output []byte, gasUsed uint64, err error, reverted bool) {
	size := len(t.orders)
	if size <= 1 {
		return
	}
	// pop call
	t.orders = t.orders[:size-1]
	if t.hooks != nil {
		t.hooks.OnExit(depth, output, gasUsed, err, reverted)
	}
}

func (t *CallStackRecorder) GetParentOrder() uint64 {
	if len(t.orders) <= 1 {
		return 0
	}
	return t.orders[len(t.orders)-2]
}
