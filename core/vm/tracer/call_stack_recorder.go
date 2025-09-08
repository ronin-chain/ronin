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

func NewTracer(hooks *tracing.Hooks) (*CallStackRecorder, *tracing.Hooks) {
	t := &CallStackRecorder{orders: make([]uint64, 1), hooks: hooks}
	wrapped := &tracing.Hooks{}
	if hooks != nil {
		wrapped = hooks.Clone()
	}
	wrapped.OnEnter = t.onEnter
	wrapped.OnExit = t.onExit
	return t, wrapped
}

func (c *CallStackRecorder) onEnter(depth int, typ byte, from common.Address, to common.Address, input []byte, gas uint64, value *big.Int, order uint64) {
	if c.hooks != nil {
		c.hooks.OnEnter(depth, typ, from, to, input, gas, value, order)
	}
	c.orders = append(c.orders, order)
}

func (c *CallStackRecorder) onExit(depth int, output []byte, gasUsed uint64, err error, reverted bool) {
	if c.hooks != nil {
		c.hooks.OnExit(depth, output, gasUsed, err, reverted)
	}
	size := len(c.orders)
	if size <= 1 {
		return
	}
	// pop call
	c.orders = c.orders[:size-1]
}

func (c *CallStackRecorder) GetParentOrder() uint64 {
	if len(c.orders) <= 1 {
		return 0
	}
	return c.orders[len(c.orders)-2]
}
