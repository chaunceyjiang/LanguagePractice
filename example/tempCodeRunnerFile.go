package main
import "fmt"
func main() {
	fmt.Println(test())
}

func test() (t int){
	t=5
	defer func(t int){
		t = t +5 
		fmt.Println(t)
	}(t)
	return  t
}