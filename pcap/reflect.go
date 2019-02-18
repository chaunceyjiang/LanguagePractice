package main

import (
	"golang.org/x/exp/errors/fmt"
	"reflect"
	"runtime"
)

func main() {

	t:=reflect.ValueOf(find).Type()
	fmt.Println(t.Name())
	if t.Kind() == reflect.Func{
		fmt.Println(runtime.FuncForPC(reflect.ValueOf(t).Pointer()).Name())
	}

}

func find()  {

}