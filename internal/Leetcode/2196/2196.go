package _2196

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	nodeMap := make(map[int]*TreeNode, len(descriptions))
	childrens := make(map[int]bool, len(descriptions))

	var root *TreeNode

	for _, description := range descriptions {
		parentValue := description[0]
		currValue := description[1]
		isLeft := description[2]

		childrens[currValue] = true

		//check that parent is already done
		parent, exist := nodeMap[parentValue]
		if !exist {
			parent = &TreeNode{Val: parentValue}
			nodeMap[parentValue] = parent
		}
		son, sonExist := nodeMap[currValue]
		if !sonExist {
			son = &TreeNode{Val: currValue}
			nodeMap[currValue] = son
		}

		if isLeft == 1 {
			parent.Left = son
		} else {
			parent.Right = son
		}
	}

	for _, description := range descriptions {
		parentValue := description[0]
		if !childrens[parentValue] {
			root = nodeMap[parentValue]
			break
		}
	}

	return root
}
