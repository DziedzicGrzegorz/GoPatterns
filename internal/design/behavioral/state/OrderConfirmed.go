package state

import "fmt"

// OrderConfirmedState represents an order that has been confirmed.
type OrderConfirmedState struct {
	order *Order
}

func (s *OrderConfirmedState) ConfirmOrder() error {
	return fmt.Errorf("order is already confirmed")
}

func (s *OrderConfirmedState) CancelOrder() error {
	fmt.Println("Order canceled")
	// Handle the order cancellation logic here
	return nil
}

func (s *OrderConfirmedState) ShipOrder() error {
	fmt.Println("Order shipped")
	s.order.setState(s.order.orderShippedState)
	return nil
}

func (s *OrderConfirmedState) DeliverOrder() error {
	return fmt.Errorf("cannot deliver the order. Order is not shipped yet")
}

// OrderShippedState represents an order that has been shipped.
type OrderShippedState struct {
	order *Order
}

func (s *OrderShippedState) ConfirmOrder() error {
	return fmt.Errorf("order is already confirmed")
}

func (s *OrderShippedState) CancelOrder() error {
	return fmt.Errorf("cannot cancel the order. Order is already shipped")
}

func (s *OrderShippedState) ShipOrder() error {
	return fmt.Errorf("order is already shipped")
}

func (s *OrderShippedState) DeliverOrder() error {
	fmt.Println("Order delivered")
	s.order.setState(s.order.orderDeliveredState)
	return nil
}
