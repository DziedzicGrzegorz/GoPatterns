package _726

import (
	"sort"
	"strconv"
	"strings"
)

func countOfAtoms(formula string) string {
	stack := make([]map[string]int, 0, len(formula)/2)
	stack = append(stack, make(map[string]int))
	n := len(formula)
	i := 0

	for i < n {
		if formula[i] == '(' {
			stack = append(stack, make(map[string]int))
			i++
		} else if formula[i] == ')' {
			currMap := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			prevMap := stack[len(stack)-1]
			i++
			start := i
			for i < n && formula[i] >= '0' && formula[i] <= '9' {
				i++
			}
			count := 1
			if start < i {
				count, _ = strconv.Atoi(formula[start:i])
			}
			for el, num := range currMap {
				prevMap[el] += num * count
			}
		} else {
			start := i
			i++
			for i < n && formula[i] >= 'a' && formula[i] <= 'z' {
				i++
			}
			element := formula[start:i]
			start = i
			for i < n && formula[i] >= '0' && formula[i] <= '9' {
				i++
			}
			count := 1
			if start < i {
				count, _ = strconv.Atoi(formula[start:i])
			}
			stack[len(stack)-1][element] += count
		}
	}

	countMap := stack[0]
	elements := make([]string, 0, len(countMap))
	for el := range countMap {
		elements = append(elements, el)
	}
	sort.Strings(elements)

	var result strings.Builder
	for _, el := range elements {
		result.WriteString(el)
		if countMap[el] > 1 {
			result.WriteString(strconv.Itoa(countMap[el]))
		}
	}

	return result.String()
}
