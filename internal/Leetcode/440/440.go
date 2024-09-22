package _440

func findKthNumber(n int, k int) int {
	current := 1
	k -= 1 // ponieważ zaczynamy od pierwszej liczby

	for k > 0 {
		steps := calculateSteps(current, current+1, n)
		if steps <= k {
			current += 1
			k -= steps
		} else {
			current *= 10
			k -= 1
		}
	}

	return current
}

func calculateSteps(curr, next, n int) int {
	steps := 0
	for curr <= n {
		steps += min(n+1, next) - curr
		curr *= 10
		next *= 10
	}
	return steps
}
