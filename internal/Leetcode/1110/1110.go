package _1110

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	//make hash set for faster lookup
	deleteSet := make(map[int]bool, len(to_delete))
	resSet := map[*TreeNode]bool{
		root: true,
	}
	//convert to_delete to hash set
	for _, v := range to_delete {
		deleteSet[v] = true
	}
	dfs(root, deleteSet, resSet)
	//return all key in resSet
	var res []*TreeNode
	for node, _ := range resSet {
		res = append(res, node)
	}

	return res
}

func dfs(root *TreeNode, deleteSet map[int]bool, resSet map[*TreeNode]bool) *TreeNode {
	res := root
	if root == nil {
		return nil
	}

	if deleteSet[root.Val] {
		res = nil
		delete(resSet, root)
		if root.Left != nil {
			resSet[root.Left] = true
		}
		if root.Right != nil {
			resSet[root.Right] = true
		}

		root.Left = dfs(root.Left, deleteSet, resSet)
		root.Right = dfs(root.Right, deleteSet, resSet)

		return res
	}

	//check left and right child
	root.Left = dfs(root.Left, deleteSet, resSet)
	root.Right = dfs(root.Right, deleteSet, resSet)

	//if current node is in deleteSet, return nil
	if deleteSet[root.Val] {
		return nil
	}

	return root
}
