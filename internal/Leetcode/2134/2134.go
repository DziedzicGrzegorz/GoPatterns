package _2134

func minSwaps(nums []int) int {
	n := len(nums)
	countOnes := 0

	for _, num := range nums {
		if num == 1 {
			countOnes++
		}
	}

	if countOnes == 0 || countOnes == n {
		return 0
	}
	currentOnes := 0
	for i := 0; i < countOnes; i++ {
		if nums[i] == 1 {
			currentOnes++
		}
	}

	minSwaps := countOnes - currentOnes

	for i := 1; i < n; i++ {
		if nums[i-1] == 1 {
			currentOnes--
		}
		if nums[(i+countOnes-1)%n] == 1 {
			currentOnes++
		}
		swaps := countOnes - currentOnes
		if swaps < minSwaps {
			minSwaps = swaps
		}
	}

	return minSwaps
}
