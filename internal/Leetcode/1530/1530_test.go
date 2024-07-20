package _1530

import (
	"testing"
)

func Test_countPairs(t *testing.T) {
	type args struct {
		root     *TreeNode
		distance int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Right: &TreeNode{
							Val: 4,
						},
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				distance: 3,
			},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 4,
						},
						Right: &TreeNode{
							Val: 5,
						},
					},
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 6,
						},
						Right: &TreeNode{
							Val: 7,
						},
					},
				},
				distance: 3,
			},
			want: 2,
		},
		{
			name: "Example 3",
			args: args{
				root: &TreeNode{
					Val: 7,
					Left: &TreeNode{
						Val: 1,
						Left: &TreeNode{
							Val: 6,
						},
					},
					Right: &TreeNode{
						Val: 4,
						Left: &TreeNode{
							Val: 5,
						},
						Right: &TreeNode{
							Val: 3,
							Left: &TreeNode{
								Val: 2,
							},
						},
					},
				},
				distance: 3,
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPairs(tt.args.root, tt.args.distance); got != tt.want {
				t.Errorf("countPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
