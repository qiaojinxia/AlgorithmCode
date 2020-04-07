package main

import "fmt"

/**
 * Created by @CaomaoBoy on 2020/3/27.
 *  email:<115882934@qq.com>
 */

type Int int
func (x Int) Less(than Data)bool{
	return x < than.(Int)
}

type UInt32 uint32
func (x UInt32) Less(than Data)bool{
	return x < than.(UInt32)
}

type String string
func (x String) Less(than Data)bool{
	return x < than.(String)
}
//红黑树 颜色
const (
	RED = true
	BLACK = false
)
type Data interface {
	Less(than Data)bool
}
//红黑树结构
type RBNode struct {
	Left   *RBNode //左节点
	Right  *RBNode //右节点
	Parent *RBNode //父节点
	Color  bool    //颜色
	Data           //key
}
type RBtree struct{
	Root *RBNode
	count uint
	NIL *RBNode
}
func less(x,y Data) bool{
	return x.Less(y)
}
func NewRbTree() *RBtree{
	node := RBNode{
		Left:   nil,
		Right:  nil,
		Parent: nil,
		Color:  BLACK,
		Data:   nil,
	}
	return &RBtree{
		Root:  &node,
		count: 0,
		NIL:   &node,
	}
}


//取得红黑树的极大值
func (rbt *RBtree) max(x *RBNode) *RBNode{
	if x == rbt.NIL{
		return rbt.NIL
	}
	for x.Right != rbt.NIL{
		x = x.Right
	}
	return x
}

//取得红黑树的极大值
func (rbt *RBtree) min(x *RBNode) *RBNode{
	if x == rbt.NIL{
		return rbt.NIL
	}
	for x.Left != rbt.NIL{
		x = x.Right
	}
	return x
}

//搜索红黑树
func (rbt *RBtree) search(x *RBNode) *RBNode{
	pnode := rbt.Root
	for pnode != rbt.NIL{
		if less(pnode.Data,x.Data){
			pnode = pnode.Right
		}else if less(x.Data,pnode.Data){
			pnode = pnode.Left
		}else{
			break
		}
	}
	return pnode
}

//取得红黑树的极小值
func (rbt *RBtree) leftRotate(x *RBNode){
	if x.Right == rbt.NIL{
		return
	}
	bak := x.Right
	bak.Left = x.Right
	if bak.Left != rbt.NIL{
		bak.Left.Parent = x
	}

}
func (rbt *RBtree) Insert(item Data) *RBNode{
	//如果 插入 nil
	if item == nil{
		return nil
	}
	node := RBNode{}
	//默认插入为 红色
	node.Color = RED
	//数据赋值
	node.Data = item
	node.Right =rbt.NIL
	node.Left =rbt.NIL
	return rbt.insert(&node)

}
func (rbt *RBtree) insert(node *RBNode) *RBNode{
	ptr := rbt.Root
	nptr := rbt.NIL
	//找到要插入的位置
	for ptr != rbt.NIL{
		nptr = ptr
		if less(node.Data,ptr.Data){
			//插入的node 比 当前节点小 往左边寻找
			ptr = ptr.Left
		}else if less(ptr.Data,node.Data){
			//插入的node 比 当前节点大 往右边寻找
			ptr = ptr.Right
		}else{
			//如果 已经插入过了
			return ptr
		}
	}
	//跟新插入节点的父节点
	node.Parent = nptr
	//如果根节点 为nil
	if nptr == rbt.NIL{
		//往根节点插入
		rbt.Root = node
	}else if less(node.Data,nptr.Data){
		//比当前节点小 往左边插入
		nptr.Left = node
	}else{
		//比当前节点大 往右边插入
		nptr.Right = node
	}

	//修复插入
	rbt.insertFixUp(node)
	rbt.count ++
	//保证 根节点 是黑色的
	rbt.Root.Color = BLACK
	return node
}
func (rbt *RBtree) InverseColor(node *RBNode ){
	if node.Color == RED{
		node.Color = BLACK
	}else{
		node.Color = RED
	}
}

//插入 后平衡 有3种情况 如果付过 父节点是黑色 那么 不需要做任何 处理
//如果 父节点 是红色 分成 2类  取决于 叔叔的颜色 如果叔叔是黑色 是一类 如果叔叔是红色 又是一类
// 如果叔叔是红色 我们需要对 父节点 叔叔节点 进行变成黑色 祖父节点 变成 红色  然后把 祖父节点当成新节点 在重复 去判断这3类情况
//如果叔叔是黑色 我们需要 我们 要判断 祖父节点 和我们之间的 关系 分为4类 如果 我们在 祖父的左左 右右 左右 右左 处理方法分别是 右旋转 坐旋转 右左旋转 左右旋转
//四种情况 旋转后 如果是 左左 或者 右右 我们只需要 反转一次 父节点 和祖父节点的颜色 如果是 右左 或者 左右 旋转 那么 我们需要反转一次 自己 和祖父的颜色
func (rbt *RBtree) insertFixUp(node *RBNode){
	//通过循环从底向上修复 修改颜色导致的失衡直到根节点
	for  node != rbt.NIL && node.Parent.Color == RED {
		//如果父节点在祖父节点的左边
		if node.Parent == node.Parent.Parent.Left {
			//如果祖父节点 在父节点的 左边 那么叔叔 一定就在右边
			uncle := node.Parent.Parent.Right
			//如果叔叔节点颜色 是红色的
			//红色的情况下 我们只需要 染色处理 将父节点 叔叔节点染成 黑色 将祖父节点染成红色
			if uncle.Color == RED{
				//父节点 黑
				node.Parent.Color = BLACK
				//叔叔节点 黑
				uncle.Color = BLACK
				//祖父节点 红
				node.Parent.Parent.Color = RED
				//移到 祖父节点重新处理 这个过程
				node = node.Parent.Parent
			}else{ //黑色的我们需要通过旋转 保持黑平衡
						//如果 叔叔节点是 黑色的 或者NIL
							if node.Parent.Left == node {
								//如果 当前节点在父节点的左左 那么 就 只想一次 右旋 父节点和当前节点颜色翻转一下
								rbt.InverseColor(node.Parent)
								rbt.InverseColor(node.Parent.Parent)
								rbt.rbRightRotate(node.Parent.Parent)
								//移到 祖父节点重新处理 这个过程
								node = node.Parent.Parent
							} else {
								//如果 当前节点在父节点的左右 那么 对 父节点先 左旋一次 再对祖先 节点再右旋一次 然后将自己和祖父节点
								//进行一次颜色反转
								rbt.InverseColor(node)
								rbt.InverseColor(node.Parent.Parent)
								rbt.rbleftRotate(node.Parent)
								rbt.rbRightRotate(node.Parent.Parent)
								//移到 祖父节点重新处理 这个过程
								node = node.Parent.Parent
							}
			}
			}else{//这里是处理 uncle 在 自己的右侧 就是比自己大的情况
				//如果祖父节点 在父节点的 右边 那么叔叔 一定就在左边
				uncle := node.Parent.Parent.Left
					if uncle.Color == RED{
						//父节点 黑
						node.Parent.Color = BLACK
						//叔叔节点 黑
						uncle.Color = BLACK
						//祖父节点 红
						node.Parent.Parent.Color = RED
						//移到 祖父节点重新处理 这个过程
						node = node.Parent.Parent
					}else{//黑色的我们需要通过旋转 保持黑平衡
						//如果 叔叔节点是 黑色的 或者NIL
						if node.Parent.Right == node {
							//如果 当前节点在父节点的右右 那么 就 只想一次 左 父节点和祖父节点惊醒一次颜色反转
							rbt.InverseColor(node.Parent)
							rbt.InverseColor(node.Parent.Parent)
							rbt.rbleftRotate(node.Parent.Parent)
							//移到 祖父节点重新处理 这个过程
							node = node.Parent.Parent
						} else {
							//如果 当前节点在父节点的右左 那么 对 父节点先 右旋一次 再对祖先 节点再左旋一次 然后将自己和祖父节点进行一次颜色反转
							//进行一次颜色反转
							rbt.InverseColor(node)
							rbt.InverseColor(node.Parent.Parent)
							rbt.rbRightRotate(node.Parent)
							rbt.rbleftRotate(node.Parent.Parent)
							//移到 祖父节点重新处理 这个过程
							node = node.Parent.Parent
						}

					}
				}
		}
}

//左旋转 和AVL树的旋转 是一样的 但是 我们旋转的同事还对parent 父节点
//的指针进行维护
func (rbt *RBtree) rbleftRotate(root *RBNode){
	//如果 根节点的右边为nil节点 那么就无法进行左旋 直接返回
	if root.Right == rbt.NIL {
		return
	}
	tmp := root.Right
	root.Right = tmp.Left
	//判断左边节点是否存在 更新 左边的节点的父节点为root
	if root.Right != rbt.NIL{
		root.Right.Parent = root
	}
	//修改旋转后的父节点
	tmp.Parent = root.Parent
	//判断 父节点的 如果旋转的是根节点 就吧根节点 设置为 旋转后的节点
	if root.Parent == rbt.NIL {
		//如果是 旋转前的root节点是root 节点下第一个节点
		//修改根节点 指向转转后的节点
		rbt.Root = tmp
	}else if root.Parent.Left == root {
		//如果 旋转前的root节点是父节点的 左边节点
		//更新成旋转后的tmp节点
		root.Parent.Left = tmp
	}else if root.Parent.Right == root{
		//如果 旋转前的root节点是父节点的 右边节点
		//更新成旋转后的tmp节点
		root.Parent.Right = tmp
	}
	//把 root 的父节点 跟新 为tmp
	root.Parent = tmp
	//tmp 的左边节点更新成 root
	tmp.Left = root
}
//左旋转 和AVL树的旋转 是一样的 但是 我们旋转的同事还对parent 父节点
//的指针进行维护
func (rbt *RBtree) rbRightRotate(root *RBNode){
	//如果 根节点的左边为nil节点 那么就无法进行右旋 直接返回
	if root.Left == rbt.NIL{
		return
	}
	tmp := root.Left
	root.Left = tmp.Right
	//下面是调整旋转后的父节点
	if root.Left != rbt.NIL{
		//子节点的 父节点设置为root节点
		root.Left.Parent = root
	}
	//跟新 tmp的父节点 为 root 的父节点
	tmp.Parent = root.Parent
	if root.Parent == rbt.NIL{
		//如果 旋转前的root节点是root 节点下第一个节点
		//修改根节点 指向转转后的节点
		rbt.Root = tmp
	}else if root.Parent.Left == root{
		//如果 旋转前的root节点是父节点的 左边节点
		//更新成旋转后的tmp节点
		tmp.Parent.Left = tmp
	}else if root.Parent.Right == root{
		//如果 旋转前的root节点是父节点的 右边节点
		//更新成旋转后的tmp节点
		tmp.Parent.Right = tmp
	}
	//将 root 的父节点指向 它的 左节点
	root.Parent = tmp
	//tmp 的右边节点更新成 root
	tmp.Right = root
}

//计算书的深度
func (rbt *RBtree)GetDepth() int{
	//声明 函数指针
	var getDeepth func(node *RBNode) int
	//内部函数 闭包调用
	getDeepth =  func(node *RBNode) int{
		if node == nil{
			return 0
		}
		if node.Left == nil && node.Right == nil{
			return 1
		}
		var leftdeep int = getDeepth(node.Left)
		var rightdeep int = getDeepth(node.Right)
		if leftdeep > rightdeep{
			return leftdeep + 1
		}else{
			return rightdeep + 1
		}
	}
	return getDeepth(rbt.Root)
}
//寻找 近似 节点 就是 性质 是一样的  具体来说 就是  左边 < root < 右边
func(rbt *RBtree) searchle (root *RBNode) *RBNode{
	ptr := rbt.Root
	ptrn := ptr
	for ptr != rbt.NIL{
		//记录指针
		ptrn = ptr
		if less(root.Data,ptr.Data){
			//如果 节点 小于
			ptr = ptr.Left
		}else if less(ptr.Data,root.Data){
			ptr = ptr.Right
		}else {//已经存在
			return ptr
		}
	}
	if less(ptrn.Data,root.Data){
		return ptrn
	}
	ptrn = rbt.desuccessor(ptrn)
	return ptrn
}
//找到 树的前驱
func (rbt *RBtree)successor(x *RBNode)*RBNode{
	if x == rbt.NIL{
		return rbt.NIL
	}
	if x.Left != rbt.NIL{
		return rbt.min(x.Right)
	}
	y := x.Parent
	for y != rbt.NIL && x == y.Right{
		x = y
		y = y.Parent
	}
	return y
}

func  (rbt *RBtree)Delete(data Data) Data{
		if data == nil{
			return nil
		}
		return rbt.delete(&RBNode{rbt.NIL,rbt.NIL,rbt.NIL,RED,data}).Data
}
func (rbt *RBtree)delete(key *RBNode)*RBNode{
	//找到 要删除的元素
	DeletedElement := rbt.search(key)
	//如果没找到 不需要删除
	if DeletedElement == rbt.NIL{
		return rbt.NIL //找不到
	}
	//返回被删除的元素
	ret := &RBNode{
		Left:   rbt.NIL,
		Right:  rbt.NIL,
		Parent: rbt.NIL,
		Color:  DeletedElement.Color,
		Data:   DeletedElement.Data,
	}
	//被删除的 节点 和 被删除节点下面的子节点
	var RealDeletedElement *RBNode
	var RealDeletedElementSub *RBNode
	// 当 树 下面 只有一个节点的情况下
	// 被删除元素
  	if DeletedElement.Left == rbt.NIL || DeletedElement.Right == rbt.NIL{
		RealDeletedElement = DeletedElement
	}else{
		//如果 实际被删除节点 节点下面有2个元素 那么 不能直接删除 找到
		//找到 他的 后继 代替它被删除
		RealDeletedElement = rbt.successor(DeletedElement)
	}
	//如果被删除节点的左边不为 nil的话
	if RealDeletedElement.Left != rbt.NIL{
		//把被删除节点的左边保存下
		RealDeletedElementSub = RealDeletedElement.Left
	}else{
		//否则 获取 被删除节点的右边 可能为 nil 可能不为 nil
		RealDeletedElementSub = RealDeletedElement.Right
	}

	//实际被删除节点的子节点的 父节点 设置为 被删除节点的父节点
	RealDeletedElementSub.Parent = RealDeletedElement.Parent

	//如果实际被删除节点的父节点 是根节点 那么 直接把子节点设置为根节点下面的节点
	if RealDeletedElement.Parent == rbt.NIL {
		rbt.Root = RealDeletedElementSub

	}else if RealDeletedElement == RealDeletedElement.Parent.Left{
		//如果 实际被删除的节点 是父节点左孩子
		//那么 把 实际被删除节点的 左孩子节点设置为 实际被删除节点的子节点
		RealDeletedElement.Parent.Left = RealDeletedElementSub
	}else{
		//如果 实际被删除节点的 是父节点的右孩子
		//那么 把实际被删除节点  右孩子 设置为 实际被删除节点的子节点
		RealDeletedElement.Parent.Right = RealDeletedElementSub
	}
	//如果实际被删除的节点 和 被删除节点 是不一样的
	// 那么这种情况 我们认为 是 左右 都不为 nil 的情况 我们通过 被删除节点的右边最小
	//的节点 代替了 被删除节点 进行 删除 所以 需要把数据交给被删除的节点,这样被删除节点不需要做什么调整
	//实际被删除节点会代替 它 被删除
	if RealDeletedElement != DeletedElement {
		DeletedElement.Data = RealDeletedElement.Data
	}
	//如果删除节点是黑色的情况下,我们需要对子节点进行修复
	//实际上所有节点删除基本上都发生在 叶子节点
	//如果 左右都不为nil的节点是 无法被实际删除的
	if RealDeletedElement.Color == BLACK && RealDeletedElementSub.Color == BLACK{
	 	rbt.deleteFixup1(RealDeletedElementSub)
	}else if RealDeletedElementSub.Color == RED{ //如果 被删除节点下面 就有 红色节点 那么我们 直接 把它变黑
		RealDeletedElementSub.Color = BLACK
	}
	//对计数--
	rbt.count --
	//返回被删除的元素
	return ret
}

//删除 修复
func (rbt *RBtree)deleteFixup1(remove *RBNode)*RBNode{

	for remove != rbt.Root && remove.Color == BLACK{
		if remove.Parent.Left == remove{//删除节点在左边
			//兄弟节点在 右边
			brother := remove.Parent.Right

			if brother.Color == RED{ //兄弟节点为红色 情况3
			   //情况 3 删除节点在左边
				brother.Color = BLACK
				remove.Parent.Color = RED
				//一次左旋
				rbt.rbleftRotate(remove.Parent)

			}else{ //为黑色 情况 1 2
				if brother.Left.Color == BLACK && brother.Right.Color == BLACK{ //情况 2 处理
					//如果 父节点是红色的情况下 我们 把父节点改为黑色
					if remove.Parent.Color == RED{
						//父节点 改为 黑色
						remove.Parent.Color = BLACK
						//循环条件结束 修复完成
						remove = rbt.Root
					}else{
						//父节点 作为删除节点继续修复
						remove = remove.Parent

					}
					//兄弟节点 改为红色
					brother.Color = RED
				}else {
					//情况1 处理方法 处理 兄弟节点 在父节点右边的情况
					//如果 父节点的右边 的右边为黑色 处理右边的左边
					if brother.Right.Color == BLACK{
						//下面旋转后新的 父节点 和原来父节点颜色 一致
						brother.Left.Color = remove.Parent.Color
						//父节点 会变成了新的兄弟节点 改为黑色
						remove.Parent.Color = BLACK
						//对兄弟节点 进行一次 右旋
						rbt.rbRightRotate(brother)
						rbt.rbleftRotate(remove.Parent)
						//等价于 break
						remove = rbt.Root

					}else{//如果 父节点的 右边为 红色处理 方法
						//下面 左旋 一次后 父节点 右边的兄弟节点 会变成新的父节点
						//所以需要继承 原先父节点的颜色
						brother.Parent.Color = remove.Parent.Color
						//父节点 会变成兄弟节点的 左子节点 设置成黑色
						remove.Parent.Color = BLACK
						//兄弟节点下面的 右节点 变黑
						brother.Right.Color = BLACK
						//准备好上面的步骤 可以进行左旋
						rbt.rbleftRotate(remove.Parent)
						//等价于 break
						remove = rbt.Root
					}
				}



			}

		}else{ //删除节点在右边
			//兄弟节点在 左边
			brother := remove.Parent.Left
			if brother.Color == RED{//兄弟节点为红色
				//情况 3 删除节点在右边
				brother.Color = BLACK
				remove.Parent.Color = RED
				//一次右旋
				rbt.rbRightRotate(remove.Parent)
			}else{//为黑色
				if brother.Left.Color == BLACK && brother.Right.Color == BLACK{//情况 2处理
					//如果 父节点是红色的情况下 我们 把父节点改为黑色
					if remove.Parent.Color == RED{
						//父节点 改为 黑色
						remove.Parent.Color = BLACK
						//循环条件结束 修复完成
						remove = rbt.Root
					}else{
						//父节点 作为删除节点继续修复
						remove = remove.Parent

					}
					//兄弟节点 改为红色
					brother.Color = RED
					}else{
					//情况1 处理方法  处理 兄弟节点 在父节点左边的情况
					//如果 父节点的左边 的左边为黑色 处理左边的右边
					if brother.Left.Color == BLACK{
						//下面旋转后新的 父节点 和原来父节点颜色 一致
						brother.Right.Color = remove.Parent.Color
						//父节点 会变成了新的兄弟节点 改为黑色
						remove.Parent.Color = BLACK
						//对兄弟节点 进行一次 右旋
						rbt.rbleftRotate(brother)
						rbt.rbRightRotate(remove.Parent)
						remove = rbt.Root
					}else{//如果 父节点的 右边为 红色处理 方法
						brother.Parent.Color = remove.Parent.Color
						remove.Parent.Color = BLACK
						brother.Left.Color = BLACK
						rbt.rbRightRotate(remove.Parent)
						remove = rbt.Root
						}
					}
				}


			}

		}

	return nil
}
func (rbt *RBtree)desuccessor(x *RBNode)*RBNode{
	if x == rbt.NIL{
		return rbt.NIL
	}
	if x.Left != rbt.NIL{
		return rbt.max(x.Left)
	}
	y := x.Parent
	for y != rbt.NIL && x == y.Left{
		x = y
		y = y.Parent
	}
	return y
}
//中序遍历
func (rbt *RBtree) InorderTraversal(){
	inordertraversal(rbt.Root,rbt.NIL)

}
func inordertraversal(root *RBNode,nil *RBNode){
	if root == nil{
		return
	}
	inordertraversal(root.Left,nil)

	if root.Data != nil{
		fmt.Println(root.Data)
	}
	inordertraversal(root.Right,nil)
}

func main(){
	rbtree := NewRbTree()
	for i:=0;i<10000;i++{
		rbtree.Insert(Int(i))
	}
	fmt.Println(rbtree.GetDepth())
	for i:=0;i<2000;i++{
		rbtree.Delete(Int(i))
	}
	fmt.Println(rbtree.GetDepth())
	for i:=10000;i<200;i++{
		rbtree.Insert(Int(i))
	}

	fmt.Println(rbtree.GetDepth())
	//for i:=1000000;i<1900000;i++{
	//	rbtree.Insert(Int(i))
	//}
	//fmt.Println(rbtree.GetDepth())
	//rbtree.InorderTraversal()
}

