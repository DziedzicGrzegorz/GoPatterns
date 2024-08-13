package _40

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0, 2^len(candidates))
	//sort array
	sort.Ints(candidates)
	dfs([]int{}, 0, 0, target, candidates, &res)

	return res
}

func dfs(cur []int, i, total, target int, candidates []int, res *[][]int) {
	if total == target {
		combination := make([]int, len(cur))
		copy(combination, cur)
		*res = append(*res, combination)
		return
	}
	if total > target || i == len(candidates) {
		return
	}
	cur = append(cur, candidates[i])
	dfs(cur, i+1, total+candidates[i], target, candidates, res)
	cur = cur[:len(cur)-1]

	for i+1 < len(candidates) && candidates[i] == candidates[i+1] {
		i++
	}
	dfs(cur, i+1, total, target, candidates, res)
}
