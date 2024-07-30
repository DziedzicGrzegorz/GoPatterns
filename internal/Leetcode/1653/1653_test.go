package _1653

import (
	"testing"
)

func Test_minimumDeletions(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test1",
			args: args{
				s: "aababbab",
			},
			want: 2,
		},
		{
			name: "Test2",
			args: args{
				s: "bbaaaaabb",
			},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDeletions(tt.args.s); got != tt.want {
				t.Errorf("minimumDeletions() = %v, want %v", got, tt.want)
			}
		})
	}
}
func BenchmarkMinDel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		minimumDeletions("aababbab")
	}
}
