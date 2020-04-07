package main

import (
	"fmt"
	"sync"
	"time"
)

/**
 * Created by @CaomaoBoy on 2020/2/18.
 *  email:<115882934@qq.com>
 */
//缓存对象
type CacheObject struct {
	Value string
	TimeToLive int64
}
//是否过期
func (c0 CacheObject) ifExpired()bool{
	if c0.TimeToLive == 0{
		return false
	}
	return time.Now().UnixNano() > c0.TimeToLive
}
type Cache struct{
	objects map[string]CacheObject //存储缓存对象
	mutex *sync.RWMutex //读写锁
	
}
func NewCache()*Cache{
	return &Cache{
		objects: make(map[string]CacheObject),
		mutex:   &sync.RWMutex{},
	}
}
func(c Cache) SetValue(cacheKey string,cachevalue string,timetolive time.Duration){
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.objects[cacheKey] =CacheObject{cachevalue,time.Now().Add(timetolive).UnixNano()}
}
func (c Cache) GetValue(cacheKey string)string{
	c.mutex.Lock()
	defer c.mutex.Unlock()
	var obj CacheObject
	obj ,ok := c.objects[cacheKey]
	if !ok{
		return ""
	}
	if obj.ifExpired() {
		delete(c.objects,cacheKey)
		return ""
	}
	return obj.Value
}

func main(){
	var c *Cache = NewCache()
	c.SetValue("caomao","xxxx",time.Second * 2)
	var name string = c.GetValue("caomao")
	fmt.Println(name)
	time.Sleep(time.Second * 2)
	name = c.GetValue("caomao")
	fmt.Println(name)

}