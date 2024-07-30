package _1653

func minimumDeletions(s string) int {
	countARight := make([]int, len(s))
	for i := len(s) - 2; i >= 0; i-- {
		countARight[i] = countARight[i+1]
		if s[i+1] == 'a' {
			countARight[i]++
		}
	}
	bCountLeft := 0
	res := len(s)
	for i, c := range s {
		deletions := bCountLeft + countARight[i]
		res = min(res, deletions)
		if c == 'b' {
			bCountLeft++
		}
	}

	return res
}
