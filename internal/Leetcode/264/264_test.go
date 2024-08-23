package _264

import (
	"testing"
)

func TestGenerateUglyNumbers(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{n: 10},
			want: 12,
		},
		{
			name: "Example 2",
			args: args{n: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateUglyNumbers(tt.args.n); got != tt.want {
				t.Errorf("GenerateUglyNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
