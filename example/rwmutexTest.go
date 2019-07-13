package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex
var count int

func main() {

	// go A()
	
	time.Sleep(1 * time.Second)
	fmt.Println("000000")
	
	mu.Lock()
	mu.Lock()

	fmt.Println("---------")
	 mu.Unlock()
	 mu.Unlock()
	count++
	fmt.Println(count)
}
// func A() {
// 	mu.RLock()
// 	fmt.Println("1111")
	
// 	B()
// }

// func B() {
// 	fmt.Println("111cccccc1")
// 	time.Sleep(5 * time.Second)
// 	fmt.Println("222222")
// 	C()
// 	fmt.Println("44555555444")
// }
// func C() {
	
// fmt.Println("44444")
// 	defer mu.RUnlock()
// }
