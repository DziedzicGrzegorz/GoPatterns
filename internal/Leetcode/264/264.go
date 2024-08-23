package _264

func GenerateUglyNumbers(n int) int {
	uglyNumbers := make([]int, n)
	uglyNumbers[0] = 1

	i2, i3, i5 := 0, 0, 0
	next2, next3, next5 := 2, 3, 5

	for i := 1; i < n; i++ {
		nextUgly := min(next2, next3, next5)
		uglyNumbers[i] = nextUgly

		if nextUgly == next2 {
			i2++
			next2 = uglyNumbers[i2] * 2
		}
		if nextUgly == next3 {
			i3++
			next3 = uglyNumbers[i3] * 3
		}
		if nextUgly == next5 {
			i5++
			next5 = uglyNumbers[i5] * 5
		}
	}

	return uglyNumbers[n-1]
}
