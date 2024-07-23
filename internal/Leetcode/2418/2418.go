package _2418

func sortPeople(names []string, heights []int) []string {
	_, sortedNames := quickSort(heights, names)
	return sortedNames
}

func quickSort(arr []int, names []string) ([]int, []string) {
	if len(arr) < 2 {
		return arr, names
	}
	left, right := 0, len(arr)-1

	// Wybór pivotu
	pivotIndex := len(arr) / 2

	// Przeniesienie pivotu na koniec
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]
	names[pivotIndex], names[right] = names[right], names[pivotIndex]

	// Partycjonowanie tablicy
	for i := range arr {
		if arr[i] > arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			names[i], names[left] = names[left], names[i]
			left++
		}
	}

	// Przeniesienie pivotu na swoje miejsce
	arr[left], arr[right] = arr[right], arr[left]
	names[left], names[right] = names[right], names[left]

	// Rekurencyjne sortowanie lewej i prawej części
	leftArr, leftNames := quickSort(arr[:left], names[:left])
	rightArr, rightNames := quickSort(arr[left+1:], names[left+1:])

	// Składanie wyników
	return append(append(leftArr, arr[left]), rightArr...), append(append(leftNames, names[left]), rightNames...)
}
