package _2196

import (
	"testing"
)

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

// TestCreateBinaryTree tests the createBinaryTree function
func TestCreateBinaryTree(t *testing.T) {
	tests := []struct {
		descriptions [][]int
		expected     *TreeNode
	}{
		{
			descriptions: [][]int{{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}, {80, 19, 1}},
			expected: &TreeNode{
				Val: 50,
				Left: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 17,
					},
				},
				Right: &TreeNode{
					Val: 80,
					Left: &TreeNode{
						Val: 19,
					},
				},
			},
		},
		{
			descriptions: [][]int{{1, 2, 1}, {2, 3, 0}, {3, 4, 1}},
			expected: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 4,
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		result := createBinaryTree(test.descriptions)
		if !isSameTree(result, test.expected) {
			t.Errorf("For descriptions %v, expected %v, but got %v", test.descriptions, test.expected, result)
		}
	}
}

// BenchmarkCreateBinaryTree benchmarks the createBinaryTree function
func BenchmarkCreateBinaryTree(b *testing.B) {
	descriptions := [][]int{{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}, {80, 19, 1}}
	for i := 0; i < b.N; i++ {
		createBinaryTree(descriptions)
	}
}
