package _2751

import (
	"slices"
)

// This leetcode was added with branch CI/benchmar (d4ba5476c40067f713ecaec63669962cccdca201) to test the performance of the solution by mistake
// cpu: 13th Gen Intel(R) Core(TM) i7-13700KF
// Benchmark_survivedRobotsHealths/length_1-24             35179855                34.85 ns/op           16 B/op          2 allocs/op
// Benchmark_survivedRobotsHealths/length_2-24             31971685                37.59 ns/op            8 B/op          1 allocs/op
// Benchmark_survivedRobotsHealths/length_4-24             15787274                76.04 ns/op           24 B/op          2 allocs/op
// Benchmark_survivedRobotsHealths/length_8-24              8495556               141.3 ns/op            56 B/op          3 allocs/op
// Benchmark_survivedRobotsHealths/length_16-24             1606058               756.0 ns/op          1042 B/op          7 allocs/op
// Benchmark_survivedRobotsHealths/length_32-24              661219              1756 ns/op            2474 B/op         11 allocs/op
// Benchmark_survivedRobotsHealths/length_64-24              304242              3793 ns/op            5682 B/op         18 allocs/op
// Benchmark_survivedRobotsHealths/length_128-24             157328              7587 ns/op           11844 B/op         26 allocs/op
// -- after add the capacity to the slice when allocating
// Benchmark_survivedRobotsHealths/length_1-24             42526818                28.52 ns/op           16 B/op          2 allocs/op
// Benchmark_survivedRobotsHealths/length_2-24             25164345                47.44 ns/op           32 B/op          2 allocs/op
// Benchmark_survivedRobotsHealths/length_4-24             16723125                72.15 ns/op           64 B/op          2 allocs/op
// Benchmark_survivedRobotsHealths/length_8-24              9208335               127.0 ns/op           128 B/op          2 allocs/op
// Benchmark_survivedRobotsHealths/length_16-24             1657602               718.9 ns/op          1178 B/op          5 allocs/op
// Benchmark_survivedRobotsHealths/length_32-24              690817              1737 ns/op            2738 B/op          8 allocs/op
// Benchmark_survivedRobotsHealths/length_64-24              315591              3755 ns/op            6201 B/op         14 allocs/op
// Benchmark_survivedRobotsHealths/length_128-24             156254              7619 ns/op           12876 B/op         21 allocs/op
func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
	indexMap := make(map[int]int)

	for i, position := range positions {
		indexMap[position] = i
	}
	slices.Sort(positions)

	stack := make([]int, 0, len(positions))

	for _, position := range positions {
		key := indexMap[position]

		if directions[key] == 'R' {
			stack = append(stack, key)
		} else {
			for len(stack) > 0 && directions[stack[len(stack)-1]] == 'R' && healths[key] > 0 {
				last := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if healths[key] > healths[last] {
					healths[last] = 0
					healths[key]--

				} else if healths[key] < healths[last] {
					healths[key] = 0
					healths[last]--
					stack = append(stack, last)

				} else {
					healths[key] = 0
					healths[last] = 0
				}

			}
			if healths[key] > 0 {
				stack = append(stack, key)
			}
		}
	}
	result := make([]int, 0, len(healths))

	for _, health := range healths {
		if health > 0 {
			result = append(result, health)
		}
	}
	return result
}
