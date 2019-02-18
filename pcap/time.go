package main

import (
	"fmt"
	"time"
)

func main() {
	var ch <-chan time.Time
	ticker := time.NewTicker(1 * time.Second)
	ch = ticker.C
	go func(ch <-chan time.Time) {
		for {

			<-ch
			fmt.Println("sdsds")
		}
	}(ch)

	time.Sleep(10 * time.Second)
	ticker.Stop()
	fmt.Println("ticker stop")
}

func add(a, b int) int {
	return a + b
}
