package _592

import (
	"testing"
)

func Test_fractionAddition(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1/3-1/2",
			args: args{
				expression: "1/3-1/2",
			},
			want: "-1/6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fractionAddition(tt.args.expression); got != tt.want {
				t.Errorf("fractionAddition() = %v, want %v", got, tt.want)
			}
		})
	}
}
