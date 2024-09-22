package _2326

type ListNode struct {
	Val  int
	Next *ListNode
}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	resultMatrix := make([][]int, m)
	for i := 0; i < m; i++ {
		resultMatrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			resultMatrix[i][j] = -1
		}
	}

	top, bottom, left, right := 0, m-1, 0, n-1

	for head != nil && top <= bottom && left <= right {
		for i := left; i <= right && head != nil; i++ {
			resultMatrix[top][i] = head.Val
			head = head.Next
		}
		top++

		for i := top; i <= bottom && head != nil; i++ {
			resultMatrix[i][right] = head.Val
			head = head.Next
		}
		right--

		if top <= bottom {
			for i := right; i >= left && head != nil; i-- {
				resultMatrix[bottom][i] = head.Val
				head = head.Next
			}
		}
		bottom--

		if left <= right {
			for i := bottom; i >= top && head != nil; i-- {
				resultMatrix[i][left] = head.Val
				head = head.Next
			}
		}
		left++
	}

	return resultMatrix
}
