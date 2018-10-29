package cache2go

import (
	"sync"
	"time"
	"log"
	"sort"
)

type CacheTable struct {
	sync.RWMutex
	name string//表名
	items map[interface{}]*CacheItem
	cleanupTimer *time.Timer//定时器
	cleanupInterval time.Duration
	loger *log.Logger

	loadData func(key interface{},args ...interface{}) *CacheItem

	addedItem func(item *CacheItem)

	aboutToDeleteItem func(item *CacheItem)
}

func (table *CacheTable)Count() int  {
	table.RLock()
	defer table.RUnlock()
	return len(table.items)
}

func (table *CacheTable)Foreach(trans func(key interface{},item *CacheItem))  {
	table.RLock()
	defer table.RUnlock()

	for k,v :=range table.items{
		trans(k,v)
	}
}

func (table *CacheTable)SetDataLoader(f func(interface{},...interface{}) *CacheItem)  {
	table.Lock()
	defer table.Unlock()

	table.loadData =f
}

func (table *CacheTable)SetAddedItemCallback(f func(item *CacheItem))  {
	table.Lock()
	defer table.Unlock()

	table.addedItem=f
}

func (table *CacheTable)SetAboutToDeleteItemCallback(f func(item *CacheItem))  {
	table.Lock()
	defer table.Unlock()
	table.aboutToDeleteItem=f
}

func (table *CacheTable)SetLoger(loger *log.Logger)  {
	table.Lock()
	defer table.Unlock()
	table.loger=loger
}
func (table *CacheTable)log(v ...interface{})  {
	if table.loger==nil{
		return
	}
	table.loger.Println(v)
}
func (table *CacheTable)expireationCheck()  {
	table.Lock()
	if table.cleanupTimer!=nil{
		table.cleanupTimer.Stop()
	}
	if table.cleanupInterval >0{
		table.log("Expiration check triggered after",table.cleanupInterval,"for table",table.name)
	}else {
		table.log("Expiration check installed for table", table.name)
	}
	now:=time.Now()

	smalllestDuration:=0*time.Second
	for key,item:=range table.items{
		item.RLock()
		lifeSpan:=item.lifeSpan
		accessedOn:=item.accessedOn
		item.RUnlock()
		if lifeSpan==0{
			continue
		}
		if now.Sub(accessedOn) >=lifeSpan{
			//table.log("\t过期删除。。。。")
			table.deleteInternal(key)//过期删除
		}else {
			if smalllestDuration == 0 || lifeSpan-now.Sub(accessedOn) <smalllestDuration{
				smalllestDuration=lifeSpan-now.Sub(accessedOn)
			}
		}
	}
	table.cleanupInterval = smalllestDuration
	if smalllestDuration>0{
		table.cleanupTimer=time.AfterFunc(smalllestDuration, func() {
			go table.expireationCheck()
		})
	}
	table.Unlock()
}

func (table *CacheTable)addInternal(item *CacheItem)  {
	table.log("Adding item with key", item.key, "and lifespan of", item.lifeSpan, "to table", table.name)
	table.items[item.key]=item
	expDur:=table.cleanupInterval
	addedItem:=table.addedItem
	table.Unlock()

	if addedItem!=nil{
		addedItem(item)
	}
	if item.lifeSpan>0 && (expDur==0 ||item.lifeSpan<expDur){
		table.expireationCheck()
	}
}

func (table *CacheTable)Add(key interface{},lifeSpan time.Duration,value interface{}) *CacheItem  {
	item:=NewCacheItem(key,lifeSpan,value)
	table.Lock()
	table.addInternal(item)
	return item
}

func (table *CacheTable)deleteInternal (key interface{}) (*CacheItem,error) {
	r,ok:=table.items[key]
	if !ok{
		return nil,ErrKeyNotFound
	}

	aboutToDeleteItem :=table.aboutToDeleteItem
	table.Unlock()

	if aboutToDeleteItem!=nil{
		aboutToDeleteItem(r)
	}

	r.RLock()
	defer r.RUnlock()

	if r.aboutToexpire!=nil{
		r.aboutToexpire(key)
	}

	table.Lock()
	table.log("Deleting item with key", key, "created on", r.createdOn, "and hit", r.accessCount, "times from table", table.name)
	delete(table.items,key)
	return r,nil
}

func (table *CacheTable)Delete(key interface{}) (*CacheItem,error) {
	table.Lock()
	defer table.Unlock()

	return table.deleteInternal(key)
}

func (table *CacheTable)Exists(key interface{})  bool {
	table.RLock()
	defer table.RUnlock()
	_,ok:=table.items[key]
	return ok
}

func (table *CacheTable)NotFoundAdd(key interface{},lifeSpan time.Duration,value interface{}) bool {
	table.Lock()
	if _,ok:=table.items[key];ok{
		table.Unlock()
		return false
	}
	item:=NewCacheItem(key,lifeSpan,value)
	table.addInternal(item)
	return true
}
func (table *CacheTable)Value(key interface{},args ...interface{}) (*CacheItem,error) {
	table.RLock()
	r,ok:=table.items[key]
	loadvalue:=table.loadData
	table.RUnlock()
	if ok{
		r.KeepAlive()
		return r,nil
	}
	if loadvalue!=nil{
		item:=loadvalue(key,args)
		if item!=nil{
			table.Add(key,item.lifeSpan,item.value)
			return item,nil
		}
		return nil,ErrKeyNotFoundOrLoadable
	}
	return nil,ErrKeyNotFound
}
func (table *CacheTable)Flush()  {
	table.Lock()
	defer table.Unlock()
	table.log("Flushing table", table.name)
	table.items=make(map[interface{}]*CacheItem)
	table.cleanupInterval=0
	if table.cleanupTimer!=nil{
		table.cleanupTimer.Stop()
	}
}

type CacheItemPair struct {
	key interface{}
	AccessCount int64
}

type CacheItemPairList []CacheItemPair

func (p CacheItemPairList)Swap(i,j int)  {
	p[i],p[j]=p[j],p[i]
}
func (p CacheItemPairList)Len() int  {
	return len(p)
}
func (p CacheItemPairList)Less(i,j int) bool {
	return p[i].AccessCount > p[j].AccessCount
}

func (table *CacheTable)MostAccessed(count int64) []*CacheItem {
	table.RLock()
	defer table.RUnlock()

	p:=make(CacheItemPairList,len(table.items))
	i:=0
	for k,v :=range table.items{
		p[i]=CacheItemPair{
			key:k,
			AccessCount:v.accessCount,
		}
		i++
	}
	sort.Sort(p)

	var r []*CacheItem
	c:=int64(0)
	for _,v:=range p{
		if c>=count{
			break
		}
		item,ok:=table.items[v.key]
		if ok{
			r=append(r,item)
		}
		c++
	}
	return r
}



