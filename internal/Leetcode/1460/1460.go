package _1460

func canBeEqual(target []int, arr []int) bool {
	if len(target) != len(arr) {
		return false
	}

	count := make(map[int]int)

	for _, num := range target {
		count[num]++
	}

	for _, num := range arr {
		if count[num] == 0 {
			return false
		}
		count[num]--
	}

	return true
}
