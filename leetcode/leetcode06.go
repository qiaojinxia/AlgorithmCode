package main

import "fmt"

type ListNode struct {
     Val int
     Next *ListNode
}

/**
 * Created by @CaomaoBoy on 2020/4/1.
 *  email:<115882934@qq.com>
 */
func reversePrint(head *ListNode) []int {
	tmp := make([]*ListNode,0)
	for head != nil{
		tmp = append(tmp,head)
		head = head.Next
	}
	tmp1 := make([]int,0)
	for i:=len(tmp)-1;i>=0;i--{
		fmt.Println(i)
		tmp1 = append(tmp1, tmp[i].Val)
	}
	return tmp1
}

func reversePrint1(head *ListNode) []int {
	if head == nil{
		return nil
	}
	tmp := make([]int,0,1)
	tmp = append(tmp,reversePrint1(head.Next)...)
	return []int{head.Val}
}
func main(){
	a := ListNode{
		Val:  1,
		Next: nil,
	}
	b := ListNode{
		Val:  2,
		Next: nil,
	}
	c := ListNode{
		Val:  3,
		Next: nil,
	}
	a.Next = &b
	b.Next = &c
	reversePrint(&a)
}

