package tracer

type CallStackRecorder struct {
	orders []uint64
}

func NewCallStackRecorder() *CallStackRecorder {
	return &CallStackRecorder{orders: make([]uint64, 1)}
}

func (c *CallStackRecorder) CaptureBegin(order uint64) {
	c.orders = append(c.orders, order)
}

func (c *CallStackRecorder) CaptureEnd() {
	size := len(c.orders)
	if size <= 1 {
		return
	}
	// pop call
	c.orders = c.orders[:size-1]
}

func (c *CallStackRecorder) GetParentOrder() uint64 {
	if c == nil || len(c.orders) < 1 {
		return 0
	}
	return c.orders[len(c.orders)-1]
}

// For testing purpose only
func (c *CallStackRecorder) GetOrders() []uint64 {
	orders := make([]uint64, len(c.orders))
	copy(orders, c.orders)
	return orders
}
