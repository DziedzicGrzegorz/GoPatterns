package _440

import (
	"testing"
)

func Test_findKthNumber(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				n: 13,
				k: 2,
			},
			want: 10,
		},
		{
			name: "Example 2",
			args: args{
				n: 10,
				k: 3,
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthNumber(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("findKthNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFindKthNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findKthNumber(4089173, 3098723)
	}
}
