package main

/* import (
	"fmt"
	"time"
)

func chanSend(ch chan []byte, b []byte) {
	b = b[:]
	b[1] = 0
	ch <- b

	fmt.Println(b)
}

func chanRecv(ch chan []byte) {
	for {
		b := <-ch
		fmt.Println(b)
	}
}
func main() {
	ch := make(chan []byte, 10)
	go chanRecv(ch)
	b := []byte{1, 2, 3, 4}
	chanSend(ch, b)
	fmt.Println(b)
	time.Sleep(3 * time.Second)
}
 */