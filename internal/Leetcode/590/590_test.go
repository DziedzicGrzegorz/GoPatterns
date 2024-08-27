package _590

import (
	"reflect"
	"testing"
)

func Test_postorder(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				root: &Node{
					Val: 1,
					Children: []*Node{
						{
							Val: 3,
							Children: []*Node{
								{
									Val:      5,
									Children: nil,
								},
								{

									Val:      6,
									Children: nil,
								},
							},
						},
					},
				},
			},
			want: []int{5, 6, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postorder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorder() = %v, want %v", got, tt.want)
			}
		})
	}
}
