package state

type Order struct {
	currentState        State
	Id                  int
	newOrderState       State
	orderConfirmedState State
	orderShippedState   State
	orderDeliveredState State
}

func NewOrder(id int) *Order {
	order := &Order{Id: id}

	// Initializing the states
	newOrderState := &NewOrderState{order: order}
	orderConfirmedState := &OrderConfirmedState{order: order}
	orderShippedState := &OrderShippedState{order: order}
	orderDeliveredState := &OrderDeliveredState{order: order}

	// Set the initial state
	order.setState(newOrderState)

	// Setting up all possible states
	order.newOrderState = newOrderState
	order.orderConfirmedState = orderConfirmedState
	order.orderShippedState = orderShippedState
	order.orderDeliveredState = orderDeliveredState

	return order
}

func (o *Order) setState(state State) {
	o.currentState = state
}

func (o *Order) ConfirmOrder() error {
	return o.currentState.ConfirmOrder()
}

func (o *Order) CancelOrder() error {
	return o.currentState.CancelOrder()
}

func (o *Order) ShipOrder() error {
	return o.currentState.ShipOrder()
}

func (o *Order) DeliverOrder() error {
	return o.currentState.DeliverOrder()
}
