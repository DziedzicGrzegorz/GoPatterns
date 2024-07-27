package _2191

import (
	"fmt"
	"sort"
)

// Funkcja sortująca tablicę nums zgodnie z mapowanymi wartościami
func sortJumbled(mapping []int, nums []int) []int {
	// Tworzymy listę par (oryginalna liczba, zmapowana liczba)
	type numPair struct {
		original int
		mapped   int
	}

	pairs := make([]numPair, len(nums))
	for i, num := range nums {
		pairs[i] = numPair{
			original: num,
			mapped:   mappingNum(num, mapping),
		}
	}

	// Sortujemy pary najpierw po mapped, a następnie po oryginalnej wartości dla tiebreaker
	sort.SliceStable(pairs, func(i, j int) bool {
		if pairs[i].mapped == pairs[j].mapped {
			return pairs[i].original == pairs[j].original
		}
		return pairs[i].mapped < pairs[j].mapped
	})

	// Tworzymy wynikową tablicę z posortowanymi oryginalnymi wartościami
	result := make([]int, len(nums))
	for i, pair := range pairs {
		result[i] = pair.original
	}

	return result
}

// Funkcja mapująca liczbę zgodnie z podaną mapą
func mappingNum(num int, mapping []int) int {
	// Obsługa przypadku gdy num jest 0
	if num == 0 {
		return mapping[0]
	}

	result := 0
	positionMultiplier := 1
	for num > 0 {
		// Pobieramy ostatnią cyfrę
		digit := num % 10
		// Przypisujemy przemapowaną cyfrę
		mappedDigit := mapping[digit]
		// Dodajemy przemapowaną cyfrę do wyniku, uwzględniając pozycję dziesiętną
		result += mappedDigit * positionMultiplier
		positionMultiplier *= 10
		// Usuwamy ostatnią cyfrę
		num /= 10
	}
	return result
}

func main() {
	mapping := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	nums := []int{990, 332, 981}
	sortedNums := sortJumbled(mapping, nums)
	fmt.Println(sortedNums) // Oczekiwany wynik: [332, 981, 990]
}
