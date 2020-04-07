package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/**
 * Created by @CaomaoBoy on 2020/2/18.
 *  email:<115882934@qq.com>
 */

type ReferenceCounter struct {
	num *uint32 //指针 计数器
	pool *sync.Pool //内存池
	removed *uint32 //删除的指针
}

//新建计数器
func NewReferenceCounter() *ReferenceCounter{
	return &ReferenceCounter{
		num:     new(uint32),
		pool:    &sync.Pool{},
		removed: new(uint32),
	}

}
func (rc *ReferenceCounter) Add(){
	atomic.AddUint32(rc.num,1)
}

func (rc *ReferenceCounter) Sub(){
	if atomic.AddUint32(rc.num,uint32(0))==0{
		atomic.AddUint32(rc.removed,1)
	}
}
func main(){
	var ref *ReferenceCounter
	ref = NewReferenceCounter()
	ref.Add()
	fmt.Println(*ref.num)
	ref.Add()
	fmt.Println(*ref.num)
}