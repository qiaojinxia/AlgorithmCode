package main

import "container/list"

/**
 * Created by @CaomaoBoy on 2020/2/22.
 *  email:<115882934@qq.com>
 */

type Stack struct {
	list *list.List
}
//新建栈
func NewStack() *Stack{
	list := list.New()
	return &Stack{list:list}
}
//压入数据
func (stack *Stack)Push(value interface{}){
	stack.list.PushBack(value)
}
//弹出站顶数据
func (stack *Stack)Pop()interface{}{
	element := stack.list.Back()
	if element != nil{
		stack.list.Remove(element)
		return element.Value
	}
	return nil
}
//查询栈顶数据
func (stack *Stack)Peak()interface{}{
	element := stack.list.Back()
	if element != nil{
		return element.Value
	}
	return nil
}
//栈的长度
func (stack *Stack)Len()int{
	return stack.list.Len()//返回长度
}
//栈是否为空
func (stack *Stack)Empty()bool{
	return stack.list.Len() == 0
}
