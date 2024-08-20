package _1937

import (
	"testing"
)

func Test_maxPoints(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int64
	}{
		{
			name:    "Example 1",
			args:    args{points: [][]int{{1, 2, 3}, {1, 5, 1}, {3, 1, 1}}},
			wantAns: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxPoints(tt.args.points); gotAns != tt.wantAns {
				t.Errorf("maxPoints() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
