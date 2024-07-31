package _1150

import (
	"math"
)

func minHeightShelves(books [][]int, shelfWidth int) int {
	cache := map[int]int{}
	return helper(0, books, shelfWidth, cache)

}

func helper(i int, books [][]int, shelfWidth int, cache map[int]int) int {
	if i == len(books) {
		return 0
	}
	if v, ok := cache[i]; ok {
		return v
	}
	curWidth := shelfWidth
	maxHeight := 0
	cache[i] = math.MaxInt64
	for j := i; j < len(books); j++ {
		width, height := books[j][0], books[j][1]
		if curWidth < width {
			break
		}
		curWidth -= width
		maxHeight = max(maxHeight, height)
		cache[i] = min(cache[i], maxHeight+helper(j+1, books, shelfWidth, cache))

	}

	return cache[i]

}
