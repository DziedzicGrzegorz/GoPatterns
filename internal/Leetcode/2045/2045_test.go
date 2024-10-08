package _2045

import (
	"testing"
)

func Test_secondMinimum(t *testing.T) {
	type args struct {
		n      int
		edges  [][]int
		time   int
		change int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1",
			args: args{
				n:      5,
				edges:  [][]int{{1, 2}, {1, 3}, {1, 4}, {3, 4}, {4, 5}},
				time:   3,
				change: 5,
			},
			want: 13,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := secondMinimum(tt.args.n, tt.args.edges, tt.args.time, tt.args.change); got != tt.want {
				t.Errorf("secondMinimum() = %v, want %v", got, tt.want)
			}
		})
	}
}

// bemch

func Benchmark_secondMinimum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		secondMinimum(5, [][]int{{1, 2}, {1, 3}, {1, 4}, {3, 4}, {4, 5}}, 3, 5)
	}
}
