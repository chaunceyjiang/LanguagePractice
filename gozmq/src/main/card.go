package main

import (
	"fmt"
)

func main() {
	var N int
	// fmt.Scanf("%d",&N)
	N = 12
	card := make([]int, N)
	card[0] = 1
	t := 0
	for i := 1; i <= N; i++ {
		for card[t] != 0 {
			t++
			t = t % N
		}
		t++
		fmt.Println(t)
		t = t % N
		for card[t] != 0 {
			t++
			t = t % N
		}
		card[t] = i
	}
	fmt.Println(card)

}
