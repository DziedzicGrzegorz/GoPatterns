package _2751

import (
	"sort"
)

func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
	indexMap := make(map[int]int)

	for i, position := range positions {
		indexMap[position] = i
	}
	sort.Ints(positions)

	stack := make([]int, 0)

	for _, position := range positions {
		key := indexMap[position]

		if directions[key] == 'R' {
			stack = append(stack, key)
		} else {
			for len(stack) > 0 && directions[stack[len(stack)-1]] == 'R' && healths[key] > 0 {
				last := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if healths[key] > healths[last] {
					healths[last] = 0
					healths[key]--

				} else if healths[key] < healths[last] {
					healths[key] = 0
					healths[last]--
					stack = append(stack, last)

				} else {
					healths[key] = 0
					healths[last] = 0
				}

			}
			if healths[key] > 0 {
				stack = append(stack, key)
			}
		}
	}
	result := make([]int, 0)
	for _, health := range healths {
		if health > 0 {
			result = append(result, health)
		}
	}
	return result
}
