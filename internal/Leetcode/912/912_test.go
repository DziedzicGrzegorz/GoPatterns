package _912

import (
	"math/rand"
	"reflect"
	"testing"
)

func Test_sortArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{nums: []int{5, 2, 3, 1}},
			want: []int{1, 2, 3, 5},
		},
		{
			name: "Example 1",
			args: args{nums: []int{5, 1, 1, 2, 0, 0}},
			want: []int{0, 0, 1, 1, 2, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArray(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

// bench
func generateRandomArray(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(size)
	}
	return arr
}

func Benchmark_sortArray(b *testing.B) {
	arr := generateRandomArray(500_000)
	for i := 0; i < b.N; i++ {
		sortArray(arr)
	}
}
