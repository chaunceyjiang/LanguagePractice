package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	for i := 0; i < len(strs)-1; i++ {
		strs[i+1] = findLCP(strs[i], strs[i+1])
	}
	return strs[len(strs)-1]
}

func findLCP(x, y string) string {
	xx := Min(len(x), len(y))
	l := 0
	for i := 0; i < xx; i++ {
		if x[i] != y[i] {
			break
		}
		l += 1
	}
	return x[:l]
}
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}
