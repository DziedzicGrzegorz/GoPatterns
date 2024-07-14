package _1701

func averageWaitingTime(customers [][]int) float64 {
	waitingTime, timeLine := 0, 0
	for _, customer := range customers {
		if timeLine < customer[0] {
			timeLine = customer[0]
		}
		waitingTime += customer[1] + (timeLine - customer[0])
		timeLine += customer[1]
	}
	return float64(waitingTime) / float64(len(customers))
}
