package _1530

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countPairs(root *TreeNode, distance int) int {
	var result int

	var dfs func(node *TreeNode) []int
	dfs = func(node *TreeNode) []int {
		if node == nil {
			return nil
		}
		if node.Left == nil && node.Right == nil {
			return []int{0}
		}
		leftDistances := dfs(node.Left)
		rightDistances := dfs(node.Right)
		for _, ld := range leftDistances {
			for _, rd := range rightDistances {
				if ld+rd+2 <= distance {
					result++
				}
			}
		}
		var distances []int
		for _, ld := range leftDistances {
			distances = append(distances, ld+1)
		}
		for _, rd := range rightDistances {
			distances = append(distances, rd+1)
		}
		return distances
	}

	dfs(root)
	return result
}
