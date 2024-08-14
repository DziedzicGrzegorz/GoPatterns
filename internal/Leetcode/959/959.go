package _959

func regionsBySlashes(grid []string) int {
	ROWS1, COLS1 := len(grid), len(grid[0])
	ROWS2, COLS2 := ROWS1*3, COLS1*3

	// Create a new grid with 3 times the size of the original grid
	newGrid := make([][]int, ROWS2)
	for i := range newGrid {
		newGrid[i] = make([]int, COLS2)
	}

	// Fill the new grid with the original grid
	for r := 0; r < ROWS1; r++ {
		for c := 0; c < COLS1; c++ {
			r2, c2 := r*3, c*3
			if grid[r][c] == '/' {
				newGrid[r2][c2+2] = 1
				newGrid[r2+1][c2+1] = 1
				newGrid[r2+2][c2] = 1
			} else if grid[r][c] == '\\' {
				newGrid[r2][c2] = 1
				newGrid[r2+1][c2+1] = 1
				newGrid[r2+2][c2+2] = 1
			}
		}
	}

	// DFS to count regions
	visited := make([][]bool, ROWS2)
	for i := range visited {
		visited[i] = make([]bool, COLS2)
	}

	count := 0
	for r := 0; r < ROWS2; r++ {
		for c := 0; c < COLS2; c++ {
			if newGrid[r][c] == 0 && !visited[r][c] {
				dfs(newGrid, visited, r, c)
				count++
			}
		}
	}
	return count
}

func dfs(grid [][]int, visited [][]bool, r, c int) {
	ROWS2, COLS2 := len(grid), len(grid[0])

	if r < 0 || c < 0 || r >= ROWS2 || c >= COLS2 || visited[r][c] || grid[r][c] != 0 {
		return
	}

	visited[r][c] = true

	neighbours := [][]int{{r + 1, c}, {r, c + 1}, {r - 1, c}, {r, c - 1}}
	for _, n := range neighbours {
		dfs(grid, visited, n[0], n[1])
	}
}
