package _624

func maxDistance(arrays [][]int) int {
	// Initialize the minimum and maximum with the first array's values
	minValue := arrays[0][0]
	maxValue := arrays[0][len(arrays[0])-1]
	maxDistance := 0

	for i := 1; i < len(arrays); i++ {
		currentArray := arrays[i]
		currentMin := currentArray[0]
		currentMax := currentArray[len(currentArray)-1]

		// Compare current array's max with global min and vice versa
		maxDistance = max(maxDistance, abs(currentMax-minValue))
		maxDistance = max(maxDistance, abs(maxValue-currentMin))

		// Update global min and max
		minValue = min(minValue, currentMin)
		maxValue = max(maxValue, currentMax)
	}

	return maxDistance
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
