package main

import (
	"fmt"
	"time"
)


func BigSlowOperation() {
	defer trace("BigSlowOperation")()()
	time.Sleep(3*time.Second)
	fmt.Println("exit")
}

func trace(msg string) func() func() {
	start:=time.Now()
	fmt.Printf("enter %s\n",msg)
	return func() func() {
		fmt.Printf("exit %s (%s)\n",msg,time.Since(start))
		return func() {
			fmt.Printf("exit %s (%s)\n",msg,time.Since(start))
		}
	}
}

func main() {
	BigSlowOperation()
}
