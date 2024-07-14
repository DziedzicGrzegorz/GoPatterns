package _1717

func maximumGain(s string, x int, y int) int {
	res := 0
	pair := func(x, y int) string {
		if x > y {
			return "ab"
		}
		return "ba"
	}(x, y)

	res += removePairs(&s, pair, max(x, y))
	res += removePairs(&s, reverseString(pair), min(x, y))

	return res
}

func removePairs(s *string, pair string, score int) int {
	res := 0
	stack := []rune{}

	for _, c := range *s {
		if c == rune(pair[1]) && len(stack) > 0 && stack[len(stack)-1] == rune(pair[0]) {
			stack = stack[:len(stack)-1]
			res += score
		} else {
			stack = append(stack, c)
		}
	}
	*s = string(stack)

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
