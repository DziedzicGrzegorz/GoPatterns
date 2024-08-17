package _1392

import (
	"reflect"
	"testing"
)

func Test_buildMatrix(t *testing.T) {
	type args struct {
		k             int
		rowConditions [][]int
		colConditions [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test 1",
			args: args{
				k: 3,
				rowConditions: [][]int{
					{1, 2},
					{2, 3},
				},
				colConditions: [][]int{
					{1, 2},
					{2, 3},
				},
			},
			want: [][]int{
				{1, 0, 0},
				{0, 2, 0},
				{0, 0, 3},
			},
		},
		{
			name: "Test 2",
			args: args{
				k: 3,
				rowConditions: [][]int{
					{1, 2},
					{2, 3},
					{3, 1},
					{2, 3},
				},
				colConditions: [][]int{
					{2, 1},
				},
			},
			want: [][]int{},
		},
		// Add more test cases if needed.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildMatrix(tt.args.k, tt.args.rowConditions, tt.args.colConditions); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_buildMatrix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buildMatrix(5, [][]int{{1, 2}, {2, 3}, {4, 5}}, [][]int{{1, 2}, {3, 4}, {4, 5}})
	}
}
