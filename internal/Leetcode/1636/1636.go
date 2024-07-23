package _1636

import (
	"slices"
)

func frequencySort(nums []int) []int {
	frequency := map[int]int{}
	for _, num := range nums {
		frequency[num]++
	}

	// Tworzenie nowej kopii nums, aby nie zmieniać oryginalnej tablicy
	sortedNums := append([]int{}, nums...)

	slices.SortFunc(sortedNums, func(a, b int) int {
		if frequency[a] == frequency[b] {
			if a > b {
				return -1
			}
			return 1
		}
		return frequency[a] - frequency[b]
	})

	return sortedNums
}
