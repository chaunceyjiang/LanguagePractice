package main

import (
	"cache2go"
	"fmt"
	"time"
)

func main() {
	cache:=cache2go.Cache("缓存2")
	cache.SetAddedItemCallback(func(item *cache2go.CacheItem) {
		fmt.Println("Added:",item.GetKey(),item.GetValue(),item.GetCreateOn())
	})

	cache.SetAboutToDeleteItemCallback(func(item *cache2go.CacheItem) {
		fmt.Println("Deleting:",item.GetKey(),item.GetValue(),item.GetCreateOn())
	})

	cache.Add("键",2*time.Second,"值")
	for i:=0;i<10;i++{
		res,err:=cache.Value("键")
		if err==nil{
			fmt.Println("Found :",res.GetKey(),"->",res.GetValue())
		}else{
			fmt.Println("Not Found :",err.Error())
		}
		time.Sleep(500*time.Millisecond)
		if i>6{
			fmt.Println("删除---------》")
			cache.Delete("键")
		}
	}


	res:=cache.Add("anotherKey",2*time.Second,"This is another test.")
	res.SetAboutExpireCallback(func(key interface{}) {
		fmt.Println("About to expire:",key.(string))
	})
	time.Sleep(6*time.Second)

}
