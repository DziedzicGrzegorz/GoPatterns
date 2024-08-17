package _1110

import (
	"sort"
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

// Funkcja, która porównuje dwie listy drzew binarnych
func compareTreeSlices(got, want []*TreeNode) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range got {
		if !isSameTree(got[i], want[i]) {
			return false
		}
	}
	return true
}

// Sortowanie drzew binarnych według wartości korzenia
func sortTreeSliceByValue(trees []*TreeNode) {
	sort.Slice(trees, func(i, j int) bool {
		return trees[i].Val < trees[j].Val
	})
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

			// Sort both slices before comparison
			sortTreeSliceByValue(got)
			sortTreeSliceByValue(tt.want)

			if !compareTreeSlices(got, tt.want) {
				t.Errorf("delNodes() = %v, want %v", got, tt.want)
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
