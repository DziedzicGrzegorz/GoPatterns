package _840

import (
	"testing"
)

func Test_numMagicSquaresInside(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				grid: [][]int{
					{4, 3, 8, 4},
					{9, 5, 1, 9},
					{2, 7, 6, 2},
				},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numMagicSquaresInside(tt.args.grid); got != tt.want {
				t.Errorf("numMagicSquaresInside() = %v, want %v", got, tt.want)
			}
		})
	}
}
