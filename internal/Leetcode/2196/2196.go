package _2196

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	nodeMap := make(map[int]*TreeNode)
	childrens := make(map[int]bool)

	//value that is in parents but not in childrens is the root
	for _, description := range descriptions {
		parentValue := description[0]
		currValue := description[1]

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

		if description[2] == 1 {
			parent.Left = son
		} else {
			parent.Right = son
		}
	}

	var root *TreeNode
	for _, desc := range descriptions {
		if childrens[desc[0]] == false {
			root = nodeMap[desc[0]]
		}
	}
	return root
}
