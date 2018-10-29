package main

// import (
// 	"fmt"
// 	"strconv"
// 	"time"
// )

// type Job struct {
// 	JobId   string
// 	JobCont string
// }

// func main() {
// 	stopWork := make(chan bool)
// 	finish := make(chan bool)
// 	defer close(finish)
// 	defer close(stopWork)
// 	var jobs, tmps []*Job
// 	for i := 0; i < 10; i++ {
// 		str := strconv.Itoa(i)
// 		job := &Job{
// 			JobCont: str,
// 			JobId:   str,
// 		}
// 		jobs = append(jobs, job)
// 		tmps = append(tmps, job)
// 	}
// 	go woker(finish, stopWork, jobs)
// 	stopMain := false
// 	timeOut := time.After(time.Duration(1) * time.Minute)
// 	quiteTimeOut := false
// 	for {
// 		if stopMain {
// 			break
// 		}
// 		select {
// 		case s := <-finish:
// 			if s == true {
// 				fmt.Printf("finish..%v", s)
// 				stopMain = true
// 				quiteTimeOut = false
// 				break
// 			}
// 		case <-timeOut:
// 			fmt.Println("You cost too much time quit now.")
// 			stopMain = true
// 			quiteTimeOut = true
// 			stopWork <- true
// 			break
// 		}

// 	}
// 	if quiteTimeOut {
// 		fmt.Println("Quit timeout ...")

// 	} else {
// 		fmt.Println("Quit common...")

// 	}

// 	size := len(tmps)
// 	for i := 0; i < size; i++ {
// 		if jobs[i] == nil {
// 			fmt.Printf("Aleady Check jobId:%s\n", i)
// 		}

// 	}
// 	time.Sleep(1 * time.Minute)
// 	fmt.Println("Quit main...")

// }
// func woker(finish chan bool, stopSignal chan bool, jobs []*Job) {
// 	size := len(jobs)
// 	stop := false
// 	rount := 1
// 	for {
// 		if stop {
// 			break
// 		}
// 		select {
// 		case stop = <-stopSignal:
// 			fmt.Println("Notity timeout stop...")
// 		default:
// 			checkNum := 0
// 			fmt.Println("Check....rount", rount)
// 			for i := 0; i < size; i++ {
// 				if jobs[i] != nil {
// 					if i%2 == 0 {
// 						jobs[i] = nil
// 						fmt.Println("Check i:", i)
// 						checkNum++
// 					} else {
// 						fmt.Println("pass check i:", i)
// 					}

// 				} else {
// 					checkNum++
// 				}
// 				if checkNum == size {
// 					stop = true
// 				}
// 			}
// 			rount++
// 			time.Sleep(2 * time.Second)
// 		}
// 	}
// 	finish <- true
// }
