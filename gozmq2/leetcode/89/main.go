package main

import (
	"fmt"
	"strconv"
)

//func grayCode(n int) []int {
//
//}

func main() {
	//fmt.Println(grayCode(2))
	//fmt.Printf("%b",2)
	// TODO
	x,err:=strconv.ParseInt(fmt.Sprintf("%b",7),2,32)
	fmt.Println(err,x)
}