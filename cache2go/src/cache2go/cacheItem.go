package cache2go

import (
	"sync"
	"time"
)

type CacheItem struct {
	sync.RWMutex

	key interface{}
	value interface{}

	lifeSpan time.Duration //生存时间

	createdOn time.Time //被创建的时间戳

	accessedOn time.Time//最后一次被访问的时间
	accessCount int64
	aboutToexpire func(key interface{}) //回调函数，移除前触发

}

func NewCacheItem(key interface{},lifeSpan time.Duration,value interface{}) *CacheItem {//工厂函数，用以生成缓存项
	t:=time.Now()
	return &CacheItem{
		key:key,
		lifeSpan:lifeSpan,
		value:value,
		createdOn:t,
		accessCount:0,
		accessedOn:t,
		aboutToexpire:nil,
	}
}

func (item *CacheItem)KeepAlive()  {
	item.Lock()
	defer item.Unlock()

	item.accessedOn=time.Now()
	item.accessCount++
}

func (item *CacheItem)GetLifeSpan() time.Duration {
	return item.lifeSpan
}
func (item *CacheItem)GetCreateOn() time.Time {
	return item.createdOn
}
func (item *CacheItem)GetKey() interface{} {
	return item.key
}
func (item *CacheItem)GetValue() interface{} {
	return item.value
}

func (item *CacheItem)SetAboutExpireCallback(f func(key interface{}))  {
	item.Lock()
	defer item.Unlock()
	item.aboutToexpire=f
}


