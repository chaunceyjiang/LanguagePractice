package main

import (
	"fmt"
	"strings"
)

var ss = "Let's take LeetCode contest"

func main() {
	fmt.Println(reverseWords(ss))
}
func reverseWords(s string) string {
	var orgStrings, oldStrings []string
	orgStrings = strings.Split(s, " ")

	for i := range orgStrings {
		oldStrings = append(oldStrings, reverseString(orgStrings[i]))
	}
	return strings.Join(oldStrings, " ")
}
func reverseString(s string) string {
	var ss []byte
	l := len(s)
	ss = make([]byte, l)
	for i := range s {
		ss[l-i-1] = s[i]
	}
	return string(ss)
}
