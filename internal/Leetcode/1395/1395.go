package _1395

func numTeams(rating []int) int {
	res := 0
	for i := 1; i < len(rating)-1; i++ {
		leftSmaller, rightLarger := 0, 0

		for j := 0; j < i; j++ {
			if rating[j] < rating[i] {
				leftSmaller++
			}
		}
		for k := i + 1; k < len(rating); k++ {
			if rating[k] > rating[i] {
				rightLarger++
			}
		}
		res += leftSmaller * rightLarger
		leftLarger := i - leftSmaller
		rightSmaller := len(rating) - i - 1 - rightLarger
		res += leftLarger * rightSmaller

	}

	return res
}
