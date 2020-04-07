package main

import (
	"fmt"
	"sync"
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

type Stack struct {
	Reference [] *ReferenceCounter //计数器
	Count int //数量
}
func (stack *Stack)New(){
	stack.Reference = make([]*ReferenceCounter,0)
}
func (stack *Stack)Push(ref *ReferenceCounter){
	stack.Reference =append(stack.Reference[:stack.Count],ref)
	stack.Count ++

}
func (stack *Stack)Pop() *ReferenceCounter{
	if stack.Count == 0{
		return nil
	}
	var length int = len(stack.Reference)
	var reference *ReferenceCounter =stack.Reference[length -1]
	if length >1 {
		stack.Reference =stack.Reference[:length-1]
	}else{
		stack.Reference =stack.Reference[0:]

	}
	stack.Count = len(stack.Reference)
	return reference
}

func main(){
	var stack *Stack = &Stack{}
	stack.New()
	var r1 *ReferenceCounter =NewReferenceCounter()
	var r2 *ReferenceCounter =NewReferenceCounter()
	var r3 *ReferenceCounter =NewReferenceCounter()
	var r4 *ReferenceCounter =NewReferenceCounter()
	var r5 *ReferenceCounter =NewReferenceCounter()
	var r6 *ReferenceCounter =NewReferenceCounter()
	stack.Push(r1)
	stack.Push(r2)
	stack.Push(r3)
	stack.Push(r4)
	stack.Push(r5)
	stack.Push(r6)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

}