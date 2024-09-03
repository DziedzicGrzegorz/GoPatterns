package state

import (
	"fmt"
)

type OrderDeliveredState struct {
	order *Order
}

func (s *OrderDeliveredState) ConfirmOrder() error {
	return fmt.Errorf("order is already delivered")
}

func (s *OrderDeliveredState) CancelOrder() error {
	return fmt.Errorf("cannot cancel the order. Order is already delivered")
}

func (s *OrderDeliveredState) ShipOrder() error {
	return fmt.Errorf("order is already delivered")
}

func (s *OrderDeliveredState) DeliverOrder() error {
	return fmt.Errorf("order is already delivered")
}
