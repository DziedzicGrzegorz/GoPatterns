package _1392

func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	// Helper function for topological sorting
	topologicalSort := func(n int, conditions [][]int) ([]int, bool) {
		adj := make([][]int, n)
		inDegree := make([]int, n)

		for _, condition := range conditions {
			u, v := condition[0]-1, condition[1]-1
			adj[u] = append(adj[u], v)
			inDegree[v]++
		}

		queue := []int{}
		for i := 0; i < n; i++ {
			if inDegree[i] == 0 {
				queue = append(queue, i)
			}
		}

		sorted := []int{}
		for len(queue) > 0 {
			node := queue[0]
			queue = queue[1:]
			sorted = append(sorted, node)
			for _, neighbor := range adj[node] {
				inDegree[neighbor]--
				if inDegree[neighbor] == 0 {
					queue = append(queue, neighbor)
				}
			}
		}

		if len(sorted) == n {
			return sorted, true
		} else {
			return nil, false
		}
	}

	// Get topological sorts
	rowOrder, rowOk := topologicalSort(k, rowConditions)
	colOrder, colOk := topologicalSort(k, colConditions)

	if !rowOk || !colOk {
		return [][]int{}
	}

	// Mapping from number to its row and column index
	rowPos := make([]int, k)
	colPos := make([]int, k)

	for i, num := range rowOrder {
		rowPos[num] = i
	}

	for i, num := range colOrder {
		colPos[num] = i
	}

	// Build the matrix
	matrix := make([][]int, k)
	for i := range matrix {
		matrix[i] = make([]int, k)
	}

	for num := 0; num < k; num++ {
		matrix[rowPos[num]][colPos[num]] = num + 1
	}

	return matrix
}
