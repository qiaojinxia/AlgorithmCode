package main

import "sync"

/**
 * Created by @CaomaoBoy on 2020/2/18.
 *  email:<115882934@qq.com>
 */
type ReferenceCounter struct {
	num *uint32 //指针 计数器
	pool *sync.Pool //内存池
	removed *uint32 //删除的指针
	weight int
}
func WeightReference() int{
	var rfs []*ReferenceCounter //数组计数器
	rfs = make([]*ReferenceCounter,10,10)
	for i:=0;i<10;i++{
		rfs[i] = new(ReferenceCounter)
	}
	var rf ReferenceCounter
	var sum int
	for _,reference := range rfs{
		sum = sum + reference.weight
	}
	return sum
}