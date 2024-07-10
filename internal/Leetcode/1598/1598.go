package _1598

func minOperations(logs []string) (depth int) {
	for _, log := range logs {
		switch log {
		case "./":
		case "../":
			if depth > 0 {
				depth--
			}
		default:
			depth++
		}
	}
	return
}
