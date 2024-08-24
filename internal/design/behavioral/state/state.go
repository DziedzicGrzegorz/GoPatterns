package state

type State interface {
	ConfirmOrder() error
	CancelOrder() error
	ShipOrder() error
	DeliverOrder() error
}
