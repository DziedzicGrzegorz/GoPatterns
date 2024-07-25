package _912

import (
	"math"
)

func sortArray(nums []int) []int {
	CountingSort(nums)
	return nums
}

func CountingSort(nums []int) []int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	// find min and max
	mi, mx := math.MaxInt64, math.MinInt64
	for _, num := range nums {
		mi = min(mi, num)
		mx = max(mx, num)
	}

	size := mx - mi + 1
	count := make([]int, size, size)
	for _, num := range nums {
		count[num-mi]++
	}

	index := 0
	for v, c := range count {
		for c > 0 {
			nums[index] = v + mi
			index++
			c--
		}
	}
	return nums
}
