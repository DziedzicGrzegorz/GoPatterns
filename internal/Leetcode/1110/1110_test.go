package _1110

import (
	"testing"
)

// Funkcja, która porównuje dwa drzewa binarne
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func Test_delNodes(t *testing.T) {
	type args struct {
		root      *TreeNode
		to_delete []int
	}
	tests := []struct {
		name string
		args args
		want []*TreeNode
	}{
		{
			name: "Example 1",
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
				to_delete: []int{3, 5},
			},
			want: []*TreeNode{
				{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 4,
						},
					},
					Right: nil,
				},
				{
					Val: 6,
				},
				{
					Val: 7,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := delNodes(tt.args.root, tt.args.to_delete)
			if len(got) != len(tt.want) {
				t.Errorf("delNodes() = %v, want %v", got, tt.want)
			}
			for i := range got {
				if !isSameTree(got[i], tt.want[i]) {
					t.Errorf("delNodes() = %v, want %v", got[i], tt.want[i])
				}
			}
		})
	}
}

func Benchmark_delNodes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		delNodes(&TreeNode{
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
		}, []int{3, 5})
	}
}
