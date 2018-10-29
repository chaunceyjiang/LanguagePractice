package main

import (
	"cache2go"
	"time"
	"fmt"
)

func main() {
	cache:=cache2go.Cache("YYYYYYYYYY")
	cache.Add("XXXXXXXXXXX",10*time.Second,1)

	res,err:=cache.Value("XXXXXXXXXXX")
	if err==nil{
		fmt.Println("Found value in cache:",res.GetValue())
	}else {
		fmt.Println("Error retrieving value from cache:", err)
	}
	// Wait for the item to expire in cache.
	time.Sleep(12 * time.Second)
	res, err = cache.Value("XXXXXXXXXXX")
	if err != nil {
		fmt.Println("Item is not cached (anymore).")
	}

	// Add another item that never expires.
	cache.Add("XXXXXXXXXXX", 0, 1)

	// Remove the item from the cache.
	cache.Delete("XXXXXXXXXXX")
	if err==nil{
		fmt.Println("Found value in cache:",res.GetValue())
	}else {
		fmt.Println("Error retrieving value from cache:", err)
	}
	for _,v:=range cache.MostAccessed(1){
		v.GetValue()
		v.GetKey()
	}
	// And wipe the entire cache table.
	cache.Flush()
}