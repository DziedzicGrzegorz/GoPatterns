package _840

import (
	"strconv"
	"strings"
)

func numMagicSquaresInside(grid [][]int) int {
	ans := 0

	for i := 0; i+2 < len(grid); i++ {
		for j := 0; j+2 < len(grid[0]); j++ {
			if grid[i][j]%2 == 0 && grid[i+1][j+1] == 5 {
				if isMagic(grid, i, j) {
					ans++
				}
			}
		}
	}

	return ans
}

func isMagic(grid [][]int, i, j int) bool {
	s := ""

	for _, num := range []int{0, 1, 2, 5, 8, 7, 6, 3} {
		s += strconv.Itoa(grid[i+num/3][j+num%3])
	}

	return strings.Contains("4381672943816729", s) || strings.Contains("9276183492761834", s)
}
