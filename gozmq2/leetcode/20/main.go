package main

import (
	"fmt"
)

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	ss := transform(s)
	m := map[string]int{
		"(": 1,
		")": -1,
		"{": 2,
		"}": -2,
		"[": 3,
		"]": -3,
	}
	stack := make([]string, 0)
	index := -1
	for i := 0; i < len(s); i++ {
		if index >= 0 && (m[stack[index]]+m[ss[i]] == 0) {
			index -= 1
			stack = stack[:len(stack)-1]
		} else {
			index += 1
			stack = append(stack, ss[i])
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func transform(s string) []string {
	ss := []byte(s)
	res := make([]string, 0)
	for i := 0; i < len(s); i++ {
		res = append(res, fmt.Sprintf("%s", ss[i:i+1]))
	}
	return res
}

func main() {
	s := "()"
	fmt.Println(isValid(s))
}
