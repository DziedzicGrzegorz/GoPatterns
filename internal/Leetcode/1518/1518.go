package _1518

func numWaterBottles(numBottles, numExchange int) int {
	emptyBottles := numBottles
	totalDrank := numBottles

	for emptyBottles >= numExchange {
		newBottles := emptyBottles / numExchange
		totalDrank += newBottles
		emptyBottles = emptyBottles%numExchange + newBottles

	}
	return totalDrank
}
