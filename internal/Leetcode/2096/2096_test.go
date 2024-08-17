package _2096

import (
	"testing"
)

func Test_getDirections(t *testing.T) {
	type args struct {
		root       *TreeNode
		startValue int
		destValue  int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 3,
						},
					},
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 6,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
				},
				startValue: 3,
				destValue:  6,
			},
			want: "UURL",
		},
		{
			name: "Test Case 2",
			args: args{
				root: &TreeNode{
					Val: 5,
					Left: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 3,
						},
					},
					Right: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 6,
						},
						Right: &TreeNode{
							Val: 4,
						},
					},
				},
				startValue: 1,
				destValue:  4,
			},
			want: "URR",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDirections(tt.args.root, tt.args.startValue, tt.args.destValue); got != tt.want {
				t.Errorf("getDirections() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_getDirections(b *testing.B) {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	for i := 0; i < b.N; i++ {
		getDirections(root, 3, 6)
	}
}
