package _2751

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_survivedRobotsHealths(t *testing.T) {
	type args struct {
		positions  []int
		healths    []int
		directions string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test case 1",
			args: args{
				positions:  []int{5, 4, 3, 2, 1},
				healths:    []int{2, 17, 9, 15, 10},
				directions: "RRRRR",
			},
			want: []int{2, 17, 9, 15, 10},
		},
		{
			name: "Test case 2",
			args: args{
				positions:  []int{3, 5, 2, 6},
				healths:    []int{10, 10, 15, 12},
				directions: "RLRL",
			},
			want: []int{14},
		},
		{
			name: "Test case 3",
			args: args{
				positions:  []int{1, 2, 5, 6},
				healths:    []int{10, 10, 11, 11},
				directions: "RLRL",
			},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := survivedRobotsHealths(tt.args.positions, tt.args.healths, tt.args.directions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("survivedRobotsHealths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func generateTestCases(limit int) []struct {
	positions  []int
	healths    []int
	directions string
} {
	var testCases []struct {
		positions  []int
		healths    []int
		directions string
	}

	for n := 1; n <= limit; n *= 2 {
		testCases = append(testCases, struct {
			positions  []int
			healths    []int
			directions string
		}{
			positions:  generateSequentialSlice(n),
			healths:    generateSequentialSlice(n),
			directions: generateDirections(n),
		})
	}

	return testCases
}

func generateSequentialSlice(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = i + 1
	}
	return slice
}

func generateDirections(n int) string {
	directions := make([]byte, n)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			directions[i] = 'R'
		} else {
			directions[i] = 'L'
		}
	}
	return string(directions)
}

func Benchmark_survivedRobotsHealths(b *testing.B) {
	limit := 128
	testCases := generateTestCases(limit)

	for _, v := range testCases {
		b.Run(fmt.Sprintf("length_%d", len(v.positions)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				survivedRobotsHealths(v.positions, v.healths, v.directions)
			}
		})
	}
}
