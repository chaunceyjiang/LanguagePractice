package cache2go

import (
	"sync"
	"os"
	"log"
)

var cache  = make(map[string]*CacheTable)
var mutex  sync.RWMutex

func Cache(tableName string) *CacheTable {
	mutex.RLock()
	t,ok:=cache[tableName]
	mutex.RUnlock()

	if !ok{
		mutex.Lock()
		t,ok = cache[tableName]
		if !ok{
			t=&CacheTable{
				name:tableName,
				items:make(map[interface{}]*CacheItem),
				loger:log.New(os.Stdout,"cache2go:",log.Ldate | log.Ltime | log.Lshortfile),
			}
			cache[tableName]=t
		}
		mutex.Unlock()
	}
	return t
}