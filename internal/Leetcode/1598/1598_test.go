package _1598

import (
	"testing"
)

func Test_minOperations(t *testing.T) {
	type args struct {
		logs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				logs: []string{"d1/", "d2/", "../", "d21/", "./"},
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				logs: []string{"d1/", "d2/", "./", "d3/", "../", "d31/"},
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.args.logs); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
