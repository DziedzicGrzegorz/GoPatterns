package _664

import (
	"math"
)

func strangePrinter(s string) int {
	n := len(s)
	mem := make([][]int, n)
	for i := range mem {
		mem[i] = make([]int, n)
	}
	return strangePrinterHelper(s, 0, n-1, mem)
}

func strangePrinterHelper(s string, i, j int, mem [][]int) int {
	if i > j {
		return 0
	}
	if mem[i][j] > 0 {
		return mem[i][j]
	}

	// Print s[i].
	mem[i][j] = strangePrinterHelper(s, i+1, j, mem) + 1

	for k := i + 1; k <= j; k++ {
		if s[k] == s[i] {
			mem[i][j] = int(math.Min(float64(mem[i][j]),
				float64(strangePrinterHelper(s, i, k-1, mem)+
					strangePrinterHelper(s, k+1, j, mem))))
		}
	}

	return mem[i][j]
}
