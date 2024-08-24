package main

import (
	"fmt"

	"github.com/dziedzicgrzegorz/goPatterns/internal/design/behavioral/state"
)

func main() {
	order := state.NewOrder(0)

	err := order.ConfirmOrder()
	if err != nil {
		fmt.Println(err)
	}

	err = order.ShipOrder()
	if err != nil {
		fmt.Println(err)
	}

	err = order.DeliverOrder()
	if err != nil {
		fmt.Println(err)
	}

	err = order.CancelOrder()
	if err != nil {
		fmt.Println(err)
	}

}
