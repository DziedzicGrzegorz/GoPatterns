package _1701

import (
	"testing"
)

func Test_averageWaitingTime(t *testing.T) {
	type args struct {
		customers [][]int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "Case 2", args: struct{ customers [][]int }{customers: [][]int{{5, 2}, {5, 4}, {10, 3}, {20, 1}}}, want: 3.25000},
		{name: "Case 1", args: struct{ customers [][]int }{customers: [][]int{{1, 2}, {2, 5}, {4, 3}}}, want: 5.00000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := averageWaitingTime(tt.args.customers); got != tt.want {
				t.Errorf("averageWaitingTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
