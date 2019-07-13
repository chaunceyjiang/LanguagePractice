package main

import "fmt"

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	index := 0
	x, y := 0, 0
	for i := 1; i <= n*n; i++ {

		res[x][y] = i

		if index == 0 {
			y++
			if y >= n || res[x][y] != 0 {
				index = 1
				y--
			}

		}
		if index == 1 {
			x++
			if x >= n || res[x][y] != 0 {
				index = 2
				x--
			}
		}
		if index == 2 {
			y--
			if y < 0 || res[x][y] != 0 {
				index = 3
				y++
			}
		}
		if index == 3 {
			x--

			if x < 0 || res[x][y] != 0 {
				index = 0
				x++
				y++

			}
		}

	}
	return res

}

func main() {
	//fmt.Println(generateMatrix(5))
	n := 5
	res := generateMatrix(n)
	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Println(res[i])
	}
}
