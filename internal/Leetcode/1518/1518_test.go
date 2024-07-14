package _1518

import (
	"testing"
)

func Test_numWaterBottles(t *testing.T) {
	type args struct {
		numBottles  int
		numExchange int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1",
			args: args{numBottles: 9, numExchange: 3},
			want: 13,
		},
		{
			name: "Test case 2",
			args: args{numBottles: 15, numExchange: 4},
			want: 19,
		},
		{
			name: "Test case 3",
			args: args{numBottles: 5, numExchange: 5},
			want: 6,
		},
		{
			name: "Test case 4",
			args: args{numBottles: 2, numExchange: 3},
			want: 2,
		},
		{
			name: "Test case 5",
			args: args{numBottles: 10, numExchange: 2},
			want: 19,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numWaterBottles(tt.args.numBottles, tt.args.numExchange); got != tt.want {
				t.Errorf("numWaterBottles() = %v, want %v", got, tt.want)
			}
		})
	}
}
