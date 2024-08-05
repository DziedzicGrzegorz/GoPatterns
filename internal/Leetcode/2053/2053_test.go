package _2053

import (
	"testing"
)

func Test_kthDistinct(t *testing.T) {
	type args struct {
		arr []string
		k   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 0",
			args: args{
				arr: []string{"d", "b", "c", "b", "c", "a"},
				k:   2,
			},
			want: "a",
		},
		{
			name: "Test case 1",
			args: args{
				arr: []string{"a", "b", "a"},
				k:   1,
			},
			want: "b",
		},
		{
			name: "Test case 2",
			args: args{
				arr: []string{"a", "b", "a"},
				k:   3,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthDistinct(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("kthDistinct() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Benchmark_KthDistinct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kthDistinct([]string{"a", "b", "a"}, 1)
	}
}
