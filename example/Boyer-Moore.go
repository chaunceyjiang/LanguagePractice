package main

// import (
// 	"fmt"
// 	"strings"
// )

// // BM 输入样例 字符串匹配算法
// const BM = `HERE IS A SIMPLE EXAMPLE`
// const ex = `EXAMPLE`

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
// func badchar(substr string, index int) int {
// 	if !strings.Contains(ex, substr) {
// 		return -1
// 	}
// 	I := strings.Index(ex, substr)
// 	if I == index {
// 		return -1
// 	}
// 	return I
// }
// func main() {
// 	BMLen := len(BM)
// 	exLen := len(ex) - 1
// 	move := 0
// 	var j int
// 	for i := exLen; i < BMLen; i += move {
// 		t := i

// 		for j = exLen; j >= 0; j-- {
// 			if ex[j] != BM[t] {
// 				move = j - badchar(string(BM[t]), j)
// 				fmt.Printf("%s %s\n", string(BM[t]), string(ex[j]))
// 				break
// 			} else {
// 				// fmt.Printf("%s\n", BM[t-j:t+1])
// 				// fmt.Printf("%s\n\n", ex[:j+1])
// 				t--
// 			}

// 		}
// 		if j == 0 {
// 			fmt.Printf("%s\n", "over")
// 			goto OVER
// 		}
// 		fmt.Println(move)
// 	}
// OVER:
// }
