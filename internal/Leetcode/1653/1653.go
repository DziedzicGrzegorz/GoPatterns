package _1653

func minimumDeletions(s string) int {
	countARight := 0
	for _, c := range s {
		if c == 'a' {
			countARight++
		}
	}
	bCountLeft := 0
	res := len(s)
	for _, c := range s {
		if c == 'a' {
			countARight--
		}
		res = min(res, bCountLeft+countARight)
		if c == 'b' {
			bCountLeft++
		}
	}

	return res
}
