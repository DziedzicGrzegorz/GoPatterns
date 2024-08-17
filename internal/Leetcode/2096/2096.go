package _2096

import (
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getDirections(root *TreeNode, startValue int, destValue int) string {
	path := []string{}

	startPath := dfs(root, path, startValue)
	destPath := dfs(root, path, destValue)
	i := 0
	for i < min(len(startPath), len(destPath)) {
		if startPath[i] != destPath[i] {
			break
		}
		i++
	}

	var res strings.Builder
	for j := i; j < len(startPath); j++ {
		res.WriteString("U")
	}
	for j := i; j < len(destPath); j++ {
		res.WriteString(destPath[j])

	}
	return res.String()

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func dfs(node *TreeNode, path []string, target int) []string {
	if node == nil {
		return nil
	}
	if node.Val == target {
		return path
	}
	//apend L to path
	path = append(path, "L")
	res := dfs(node.Left, path, target)
	if res != nil {
		return res
	}
	//pop L from path
	path = path[:len(path)-1]
	//apend R to path
	path = append(path, "R")
	res = dfs(node.Right, path, target)
	if res != nil {
		return res
	}
	//pop R from path
	path = path[:len(path)-1]

	return nil
}
