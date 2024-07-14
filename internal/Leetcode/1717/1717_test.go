package _1717

import (
	"testing"
)

func Test_maximumGain(t *testing.T) {
	type args struct {
		s string
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1",
			args: args{
				s: "cdbcbbaaabab",
				x: 4,
				y: 5,
			},
			want: 19,
		},
		{
			name: "Test case 2",
			args: args{
				s: "aabbaaxybbaabb",
				x: 5,
				y: 4,
			},
			want: 20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumGain(tt.args.s, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("maximumGain() = %v, want %v", got, tt.want)
			}
		})
	}
}
