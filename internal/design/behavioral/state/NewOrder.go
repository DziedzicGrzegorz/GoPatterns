package state

import "fmt"

// NewOrderState represents a new order that has been placed.
type NewOrderState struct {
	order *Order
}

func (s *NewOrderState) ConfirmOrder() error {
	fmt.Println("Order confirmed")
	s.order.setState(s.order.orderConfirmedState)
	return nil
}

func (s *NewOrderState) CancelOrder() error {
	fmt.Println("Order canceled")
	// Handle the order cancellation logic here
	return nil
}

func (s *NewOrderState) ShipOrder() error {
	return fmt.Errorf("cannot ship the order. Order is not confirmed yet")
}

func (s *NewOrderState) DeliverOrder() error {
	return fmt.Errorf("cannot deliver the order. Order is not shipped yet")
}
