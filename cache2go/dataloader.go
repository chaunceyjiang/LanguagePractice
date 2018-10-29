package main

import (
	"cache2go"
	"strconv"
	"fmt"
)

func main() {
	cache:=cache2go.Cache("缓存")
	cache.SetDataLoader(func(i interface{}, i2 ...interface{}) *cache2go.CacheItem {
		val:="This is a test with key "+i.(string)
		item:=cache2go.NewCacheItem(i,0,val)
		return item
	})
	for i:=0;i<10;i++{
		res,err:=cache.Value("SomeKey_",strconv.Itoa(i))
		if err==nil{
			fmt.Println("Found value in cache:",res.GetKey(),res.GetValue())
		}else {
			fmt.Println("Error retrieving value from cache:",err)
		}
	}
}
