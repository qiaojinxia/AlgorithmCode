package main

/**
 * Created by @CaomaoBoy on 2020/3/3.
 *  email:<115882934@qq.com>
 */

type Map struct {
	Level int  //你的等级
	Left  *Map //左边的节点
	Right *Map //右边的节点
}
type BinaryTree struct {
	Root *Map //初始村
	Size int  //打怪点数量
}

func NewBinaryTree() *BinaryTree{
	bst := &BinaryTree{}
	bst.Size = 0
	bst.Root = nil
	return bst
}
//获取刷怪点个数
func(bst *BinaryTree) GetSize() int{
	return bst.Size
}

//有没有刷怪点
func(bst *BinaryTree) IsEmpty() bool{
	return bst.Size == 0
}

//把你加进地图
func (bst *BinaryTree)Add(data int) {
	if bst.Root == nil{
		bst.Root = &Map{data,nil,nil}
	}else{
		bst.add(bst.Root,data)
	}
}

//内部方法 给定地图,把你 加入节点
func (bst *BinaryTree) add(n *Map,data int) *Map {
	//如果 没有找到任何一张可以刷怪的地图
	if n == nil{
		//想像一下 你往左走 往右走 走到尽头了
		//找不到合适的刷怪的地图 那么 就 只能 自己创建一张啦
		//地图数量 加1张
		bst.Size ++
		//如果都没又找到 合适的 那你只能 自己创建一张了 返回给上一层递归调用
		return &Map{data,nil,nil}
	}else{
		//你的等级 比较小 往左走
		if data < n.Level {
			//更新 成左边的一张地图 再 调用 add 继续走
			n.Left = bst.add(n.Left,data)
		//你的等级 比较高 往右走
		}else if data > n.Level {
			//更新 成右边的一张地图 再 调用 add 继续走
			n.Right = bst.add(n.Right,data)
		}
	}
	return n
}

//这个是对外部暴露的,只需要把你扔进去就好啦
func (bst *BinaryTree) IsIn(data int){
	//调用自己内部的方法 把自己存储的地图 传入
	//注意IsIn 可以被其他包调用 可以理解为java 的public 而 下面那个方法 isIn 是内部方法  private
	bst.isIn(bst.Root,data)
}

//寻找有没有适合你的刷怪点 每次都是一样的 往右 或往左 所以采用递归
func (bst *BinaryTree) isIn(node *Map,data int) bool{
	//如果走到底 没有合适的打怪区 就返回fasle
	if node == nil {
		return false
	//如果 这块地方 刷怪等级太高了 就往左走
	}else if node.Level > data{
		//由于走到下一章地图 这个过程是重复的 所以调用递归 node.Left 把地图改成左边的地图
		bst.isIn(node.Left,data)

	}else if node.Level < data{
		//由于走到下一章地图 这个过程是重复的 所以调用递归 node.Left 把地图改成右边的地图
		bst.isIn(node.Right,data)
	}
	// 如果 不是大于 小于 那么只能是 等于喽  如果 你等等级 和地图等级一样说明可以找到合适的刷怪点
	return true
}
//找到地图里面的最大 这里就很容易理解了 一直往右走 必然是 最大的 所以从
//根节点 一直往右走 就好啦
func (bst *BinaryTree) FindMax() *Map {
	return bst.findmax(bst.Root)
}

func (bst *BinaryTree) findmax(node *Map) *Map {
	//如果走到了 最后 就 返回 递归的结束条件
	//一般递归中 必须要 包含 递归结束的基条件
	if node.Right == nil{
		return node
	}else{
		//递归 当 上面那步 不能往右走了 结果会返回 函数就结束啦
		return bst.findmax(node.Right)
	}
}

//min 如上
func (bst *BinaryTree) FindMin() *Map {
	return bst.findmin(bst.Root)
}

func (bst *BinaryTree) findmin(node *Map) *Map {
	if node.Left == nil{
		return node
	}else{
		//递归 当 上面那步 不能往右走了 结果会返回 函数就结束啦
		return bst.findmin(node.Left)
	}
}