package _1508

import (
	"sort"
)

func rangeSum(nums []int, n int, left int, right int) int {
	const MOD int = 1e9 + 7

	// Generate all subarray sums
	subarraySums := []int{}
	for i := 0; i < n; i++ {
		currentSum := 0
		for j := i; j < n; j++ {
			currentSum += nums[j]
			subarraySums = append(subarraySums, currentSum)
		}
	}

	// Sort the subarray sums
	sort.Ints(subarraySums)

	// Compute the sum of the specified range
	sum := 0
	for i := left - 1; i < right; i++ { // 1-based index, so subtract 1 for 0-based array
		sum = (sum + subarraySums[i]) % MOD
	}

	return sum
}
