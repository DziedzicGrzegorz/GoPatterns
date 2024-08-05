package _2053

func kthDistinct(arr []string, k int) string {
	mp := make(map[string]int)
	for _, val := range arr {
		mp[val] += 1
	}
	ku := 1

	for _, val := range arr {
		if mp[val] == 1 && k == ku {
			return val
		} else if mp[val] == 1 {
			ku++
		}
	}
	return ""
}
