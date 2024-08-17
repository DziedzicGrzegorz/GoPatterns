package _2678

func countSeniors(details []string) int {
	// extract age from string
	// convert age to int
	// count > 60
	seniorCount := 0
	for _, eachPassenger := range details {
		// 0:10 phone numebr 10:10 gender 11:13 age
		extractedAge := eachPassenger[11:13]
		age := 10*(extractedAge[0]-'0') + extractedAge[1] - '0'
		if age > 60 {
			seniorCount++
		}
	}
	return seniorCount
}
