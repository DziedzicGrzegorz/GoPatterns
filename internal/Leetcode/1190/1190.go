package _1190

import (
	"strings"
)

func reverseParentheses(s string) string {
	stack := []string{}

	for _, c := range s {
		if c == ')' {
			portion := []string{}
			for stack[len(stack)-1] != "(" {
				portion = append(portion, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			stack = append(stack, portion...)
		} else {
			stack = append(stack, string(c))
		}
	}
	return strings.Join(stack, "")
}
