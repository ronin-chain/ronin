package live

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/eth/tracers"
)

func init() {
	tracers.LiveDirectory.Register("callTracerOrder", newCallTracerOrder)
}

type callOrderFrame struct {
	Order uint64
}

type callTracerOrder struct {
	noop
	callstack []callOrderFrame
}

// newCallTracerOrder returns a live tracer which tracks order and parent order of a tx
func newCallTracerOrder(_ json.RawMessage) (*tracing.Hooks, error) {
	t := &callTracerOrder{callstack: make([]callOrderFrame, 1)}
	return &tracing.Hooks{
		OnEnter:        t.onEnter,
		OnExit:         t.onExit,
		GetParentOrder: t.getParentOrder,
	}, nil
}

// onEnter is called when EVM enters a new scope (via call, create or selfdestruct).
func (t *callTracerOrder) onEnter(depth int, typ byte, from common.Address, to common.Address, input []byte, gas uint64, value *big.Int, order uint64) {
	call := callOrderFrame{
		Order: order,
	}
	t.callstack = append(t.callstack, call)
}

// onExit is called when EVM exits a scope, even if the scope didn't
// execute any code.
func (t *callTracerOrder) onExit(depth int, output []byte, gasUsed uint64, err error, reverted bool) {
	size := len(t.callstack)
	if size <= 1 {
		return
	}

	// pop call
	t.callstack = t.callstack[:size-1]
}

func (t *callTracerOrder) getParentOrder() uint64 {
	if len(t.callstack) <= 1 {
		return 0
	}
	return t.callstack[len(t.callstack)-2].Order
}
