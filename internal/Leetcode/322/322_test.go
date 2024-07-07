package _322

import (
	"testing"
)

func TestCoinChange(t *testing.T) {
	tests := []struct {
		coins  []int
		amount int
		expect int
	}{
		{[]int{1, 2, 5}, 11, 3},
		{[]int{2}, 3, -1},
		{[]int{1}, 0, 0},
		{[]int{1, 2, 5}, 0, 0},
		{[]int{1, 2, 5}, 7, 2},
		{[]int{1, 5, 10}, 18, 5},
		{[]int{2, 5, 10}, 27, 4},
	}

	for _, test := range tests {
		result := coinChange(test.coins, test.amount)
		if result != test.expect {
			t.Errorf("For coins %v and amount %d, expected %d but got %d", test.coins, test.amount, test.expect, result)
		}
	}
}
