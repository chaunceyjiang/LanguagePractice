package main

import "fmt"

func maxProfit(prices []int) int {
	total := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > total {
				total = prices[j] - prices[i]
			}
		}
	}
	return total
}

func main() {
	var prices = []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
