package main

import (
	"fmt"
)

var orgString = "A man, a plan, a canal: Panama"

func main() {
	fmt.Println(reverseString(orgString))
}

func reverseString(s string) string {
	var ss []byte
	l:=len(s)
	ss = make([]byte,l)
	for i:=range s{
		ss[l-i-1] = s[i]
	}
	return string(ss)
}
