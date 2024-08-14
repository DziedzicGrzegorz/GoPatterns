package _2134

import (
	"testing"
)

func Test_minSwaps(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case1",
			args{
				[]int{1, 0, 1, 0, 1},
			},
			1,
		},
		{
			"case2",
			args{
				[]int{0, 1, 1, 1, 0, 0, 1, 1, 0},
			},
			2,
		},
		{
			"case3",
			args{
				[]int{1, 1, 0, 0, 1},
			},
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSwaps(tt.args.nums); got != tt.want {
				t.Errorf("minSwaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
