package _860

func lemonadeChange(bills []int) bool {
	count5, count10 := 0, 0

	for _, bill := range bills {
		if bill == 5 {
			count5++
		} else if bill == 10 {
			if count5 > 0 {
				count5--
				count10++
			} else {
				return false // Not enough $5 bills to give change
			}
		} else { // bill == 20
			if count10 > 0 && count5 > 0 {
				count10--
				count5--
			} else if count5 >= 3 {
				count5 -= 3
			} else {
				return false // Not enough change to give for $20 bill
			}
		}
	}

	return true
}
