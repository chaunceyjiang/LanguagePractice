package main

// import "fmt"

//import (
//	"log"
//	"sync"
//	"time"
//	"math/rand"
//)
//
//func main() {
//	wg:=sync.WaitGroup{}
//	for i:=0;i<10;i++{
//		wg.Add(1)
//		go func(wgs *sync.WaitGroup) {
//			defer wgs.Done()
//			time.Sleep(time.Second * time.Duration(rand.Intn(5)))
//			log.Println("Success!")
//		}(&wg)
//	}
//	wg.Wait()
//}
//

//type Test struct {
//	Name string
//}
//
//
//func main()  {
//	var  list map[string]*Test
//	list=make(map[string]*Test)
//	name:=Test{
//		Name:"xsdsxx",
//	}
//	list["dada"]=&name
//	fmt.Println(list["dada"].Name)
//}