package main

import "errors"

/**
 * Created by @CaomaoBoy on 2020/3/6.
 *  email:<115882934@qq.com>
 */
type AVLnode struct {
	Data interface{} //数据
	Left *AVLnode //左边的节点
	Right *AVLnode //右边的节点
	height int //高度
}
//定义比较大小函数指针
type comparator func (a,b interface{})int
//比较大小
func Max(a,b int) int{
	if a >= b{
		return a
	}else{
		return b
	}
}

//往左 不断递归掉用 得到最小值
func (avlnode *AVLnode) FindMin() *AVLnode{
	var finded *AVLnode
	if avlnode.Left != nil{
		//递归 查找最小
		finded = avlnode.Left.FindMin()
	}else{
		finded = avlnode
	}
	return finded
}
//往右 递归调用 得到最大值
func (avlnode *AVLnode) FindMax() *AVLnode{
	var finded *AVLnode
	if avlnode.Right != nil{
		//递归 查找最小
		finded = avlnode.Left.FindMax()
	}else{
		finded = avlnode
	}
	return finded
}
var compare comparator

func (avlnode *AVLnode) Find(data interface{}) *AVLnode{
	var finded *AVLnode = nil
	switch compare(data,avlnode.Data){
	case -1:
		finded = avlnode.Left.Find(data)
	case 1:
		finded = avlnode.Right.Find(data)
	case 0:
		return avlnode

	}
	return finded
}

//新建Node节点
func NewNode(data interface{}) *AVLnode{
	node := new(AVLnode)
	node.Data = data
	node.Left = nil
	node.Right = nil
	node.height = 0
	return node
}
//新建 AVL tree
func NewAVLTree(data interface{},myfunc comparator) (*AVLnode,error){
	if data == nil && myfunc == nil{
		return nil,errors.New("nil error")
	}
	compare = myfunc
	return NewNode(data),nil
}

//获取数据
func (avlnode *AVLnode) Getdata() interface{}{
	return avlnode.Data
}
//设置数据
func (avlnode *AVLnode) Setdata(data interface{}){
	avlnode.Data = data
}
//获取左边节点
func (avlnode *AVLnode) GetLeft() *AVLnode{
	return avlnode.Left
}
//获取高度
func (avlnode *AVLnode) GetHeight() int{
	return avlnode.height
}
//获取右边节点
func (avlnode *AVLnode) GetRight() *AVLnode{
	return avlnode.Right
}
//右旋
func (avlnode *AVLnode) RightRotate() *AVLnode{
	rgnode := avlnode.Left
	avlnode.Left = rgnode.Right
	rgnode.Right = avlnode
	//求深度
	avlnode.height =Max(avlnode.Left.GetHeight(),avlnode.Right.GetHeight()) + 1
	rgnode.height = Max(avlnode.Left.GetHeight(),avlnode.Right.GetHeight()) + 1
	return rgnode
}

//左右旋
func (avlnode *AVLnode) LRightRotate() *AVLnode{
	avlnode.Left = avlnode.Left.LeftRotate()
	return avlnode.RightRotate()
}

//左旋
func (avlnode *AVLnode) LeftRotate() *AVLnode{
	lfnode := avlnode.Right
	lfnode.Right = lfnode.Left
	lfnode.Left = avlnode
	//求深度
	avlnode.height =Max(avlnode.Left.GetHeight(),avlnode.Right.GetHeight()) + 1
	lfnode.height = Max(avlnode.Left.GetHeight(),avlnode.Right.GetHeight()) + 1
	return lfnode
	
}
//右左旋
func (avlnode *AVLnode) LReftRotate() *AVLnode{
	avlnode.Right = avlnode.Right.RightRotate()
	return avlnode.LeftRotate()
}
//负载因子 差距为 2 不平衡
func (avlnode *AVLnode) adjust() *AVLnode{
	if avlnode.Right.GetHeight() - avlnode.Left.GetHeight() == 2 {

	}else if avlnode.Left.GetHeight() - avlnode.Right.GetHeight() ==2  {


	}
	return nil
}

//计算负载 因子

