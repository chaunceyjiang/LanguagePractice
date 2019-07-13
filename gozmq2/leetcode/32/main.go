package main

import "fmt"

func transform(s string) []string {
	ss := []byte(s)
	res := make([]string, 0)
	for i := 0; i < len(s); i++ {
		res = append(res, fmt.Sprintf("%s", ss[i:i+1]))
	}
	return res
}

func longestValidParentheses(s string) int {
	ss := transform(s)
	stack := make([]int, 0)
	max := -1
	stack = append(stack, -1)
	for i := 0; i < len(ss); i++ {
		if ss[i] == "(" {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				max = Max(i-stack[len(stack)-1], max)
			}
		}
	}
	return max
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestValidParentheses("))))()())"))
}
