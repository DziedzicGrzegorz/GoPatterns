package _959

import (
	"testing"
)

func Test_regionsBySlashes(t *testing.T) {
	type args struct {
		grid []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				grid: []string{" /", "/ "},
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				grid: []string{" /", "  "},
			},
			want: 1,
		},
		{
			name: "Example 3",
			args: args{
				grid: []string{"\\\\/", "/\\\\"},
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := regionsBySlashes(tt.args.grid); got != tt.want {
				t.Errorf("regionsBySlashes() = %v, want %v", got, tt.want)
			}
		})
	}
}
