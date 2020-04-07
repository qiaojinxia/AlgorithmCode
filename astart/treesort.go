package main

import (
	"bytes"
	"container/list"
	"fmt"
	"strconv"
)

/**
 * Created by @CaomaoBoy on 2020/3/4.
 *  email:<115882934@qq.com>
 */



//中序 遍历
func(bst *BinaryTree) InOrder(){
	fmt.Print("中序遍历:")
	bst.inorder(bst.Root)
	fmt.Println()
}
func(bst *BinaryTree) inorder(node *Map){
	if node == nil{
		return
	}
	bst.inorder(node.Left)
	fmt.Print(node.Level," ")
	bst.inorder(node.Right)


}

//前序遍历
func(bst *BinaryTree) PreOrder(){
	fmt.Print("前序遍历:")
	bst.preorder(bst.Root)
	fmt.Println()

}
func(bst *BinaryTree) preorder(node *Map){
	if node == nil{
		return
	}
	fmt.Print(node.Level," ")
	bst.preorder(node.Left)
	bst.preorder(node.Right)


}

//后序遍历
func(bst *BinaryTree) PostOrder(){
	fmt.Print("后序遍历:")
	bst.postorder(bst.Root)
	fmt.Println()
}
func(bst *BinaryTree) postorder(node *Map){
	if node == nil{
		return
	}
	bst.postorder(node.Left)
	bst.postorder(node.Right)
	fmt.Print(node.Level," ")

}

func (bst *BinaryTree) String() string{
	var buffer bytes.Buffer
	bst.GenerateBSTstring(bst.Root,0,&buffer)
	return buffer.String()
}
func (bst *BinaryTree) GenerateBSTstring(node *Map,depth int,buffer *bytes.Buffer){
	if node == nil{
		buffer.WriteString(bst.GenerateDepthstring(depth) + "nil\n")//空节点
		return
	}
	buffer.WriteString(bst.GenerateDepthstring(depth) + strconv.Itoa(node.Level) + "\n")
	bst.GenerateBSTstring(node.Left,depth +1,buffer)
	bst.GenerateBSTstring(node.Right,depth +1,buffer)

}
func (bst *BinaryTree) GenerateDepthstring(depth int) string{
	var buffer bytes.Buffer
	for i:=0;i<depth;i++{
		buffer.WriteString("--")
	}
	return buffer.String()
}


//删除最大
func (bst *BinaryTree) removemax(node *Map) *Map{
	if node.Right == nil{
		bst.Size --
		//备份左孩子
		leftbak := node.Left
		return leftbak
	}
	node.Right = bst.removemax(node.Right)
	return node
}


//删除最大
func (bst *BinaryTree) RemoveMax() *Map{
	max := bst.findmax(bst.Root)
	bst.Root = bst.removemax(bst.Root)
	return max
}
//删除最小
func (bst *BinaryTree) RemoveMin() *Map{
	min := bst.findmin(bst.Root)
	bst.Root = bst.removemin(bst.Root)
	return min
}

//删除最小
func (bst *BinaryTree) removemin(node *Map) *Map{
	if node.Left == nil{
		//备份一下要被删除的左孩子 它的右孩子
		rightNode := node.Right
		bst.Size --
		return rightNode
	}
	node.Left = bst.removemin(node.Left)
	return node
}

func (bst *BinaryTree) Remove(data int)  {
	if bst.Root == nil || bst.Size == 0{
		panic("tree empty")
	}
	 bst.Root = bst.remove(bst.Root,data)
}

func (bst *BinaryTree) remove(node *Map,data int) *Map{
	if node == nil{
		return nil
	}
	if data < node.Level{
		node.Left = bst.remove(node.Left,data)
		return node
		
	}else if data > node.Level {
		node.Right = bst.remove(node.Right,data)
		return node

	}else{
		//上面2 个else 不满足的话 这里一定就是 node.Level == data
		//如果左节点 为 空的话  把 右孩子备份 返回给父节点
		//注意 递归调用 就好像 剥洋葱一样  将原来的一层一层 通过一次一次
		//递归来拆分 修改完后 return 每一层 组合起来

		//处理左边为空 和删除最大和最小一个原理
		if node.Left == nil{
			//取出 右边的节点
			rightbak := node.Right
			//nil解出引用 对功能不影响
			node.Right = nil
			bst.Size --
			return rightbak
		}
		//处理 右边为空 和删除最大和最小一个原理
		if node.Right == nil{
			leftbak := node.Left
			//nil解出引用 对功能不影响
			node.Left = nil
			bst.Size --
			return leftbak
		}
		//找到最大的左孩子
		rpnode := bst.findmax(node.Left)
		//删除最大也就是上面找到的那么节点rpnode 后返回左子树
		rpnode.Left = bst.removemax(node.Left)
		rpnode.Right = node.Right
		//nil解出引用 对功能不影响
		node.Left = nil
		node.Right = nil
		return rpnode

	}

}

//非递归 中序遍历
func (bst *BinaryTree)InOrderNoRecursion() []int{
	mybst := bst.Root
	mystack := list.New()
	res := make([]int,0)
	for mybst != nil || mystack.Len() !=0{
		for mybst != nil{
			mystack.PushBack(mybst)
			mybst = mybst.Left
		}
		if mystack.Len() != 0{
			v := mystack.Back()
			mybst = v.Value.(*Map)
			fmt.Println(mybst.Level)
			res = append(res,mybst.Level)
			mybst = mybst.Right
			mystack.Remove(v)
		}
	}

	return res

}
func (bst *BinaryTree) StringNoRecursion() string{
	var buffer bytes.Buffer
	bst.GenerateBSTstring(bst.Root,0,&buffer)
	return buffer.String()
}
func (bst *BinaryTree) GenerateBSTstringNoRecursion(node *Map,depth int,buffer *bytes.Buffer){
	if node == nil{
		buffer.WriteString(bst.GenerateDepthstring(depth) + "nil\n")//空节点
		return
	}
	//创建一个 副本来保存当前节点
	mybst := bst.Root
	//新建栈
	mystack := list.New()
	for mystack.Len() != 0 || mybst != nil{
		//如果节点不为空 就往左 压栈直到 最后一个子节点
		for mybst != nil {
			depth = mystack.Len()
			buffer.WriteString(bst.GenerateDepthstring(depth) + strconv.Itoa(node.Level) + "\n")
			//节点压栈
			mystack.PushBack(mybst)
			//更新成左节点
			mybst = mybst.Left
			depth += 1
		}
		//上面一步 如果有压栈数据
		if  mystack.Len()!= 0{
			//从栈里取出节点
			tmp := mystack.Back()
			//取出 *map 类型数值
			lbst := tmp.Value.(*Map)
			//将 指针移到节点 取出节点的右边
			mybst = lbst.Right
			//删除当前节点
			mystack.Remove(tmp)
		}
	}


}
//非递归 前序遍历
func (bst *BinaryTree)PreOrderNoRecursion() []int{
	//创建一个 副本来保存当前节点
	mybst := bst.Root
	//新建栈
	mystack := list.New()
	//保存遍历的结果
	res := make([]int,0)
	for mystack.Len() != 0 || mybst != nil{
		//如果节点不为空 就往左 压栈直到 最后一个子节点
		for mybst != nil {
			//把结果保存
			res = append(res, mybst.Level)
			fmt.Println(mystack.Len())
			//节点压栈
			mystack.PushBack(mybst)
			//更新成左节点
			mybst = mybst.Left

		}
		//上面一步 如果有压栈数据
		if  mystack.Len()!= 0{
			//从栈里取出节点
			tmp := mystack.Back()
			//取出 *map 类型数值
			lbst := tmp.Value.(*Map)
			//将 指针移到节点 取出节点的右边
			mybst = lbst.Right
			//删除当前节点
			mystack.Remove(tmp)
		}
	}

	return res

}



//非递归 后序遍历
func (bst *BinaryTree)PostOrderNoRecursion() []int{
	//创建一个 副本来保存当前节点
	mybst := bst.Root
	//新建栈
	mystack := list.New()
	//保存遍历的结果
	res := make([]int,0)
	//提前被封
	var PreVisited *Map
	for mystack.Len() != 0 || mybst != nil{
		//如果节点不为空 就往左 压栈直到 最后一个子节点
		for mybst != nil {
			//节点压栈
			mystack.PushBack(mybst)
			//更新成左节点
			mybst = mybst.Left
		}
		v := mystack.Back()//取出栈中的节点
		top := v.Value.(*Map)
		//[12 31 27 41 36 59 63 60 58 57 78 70 69 97 84 65]
		if (top.Left == nil && top.Right == nil) || top.Right == nil ||(PreVisited == top.Right) {
			res= append(res,top.Level)
			PreVisited = top
			mystack.Remove(v)
		}else{
			mybst = top.Right
		}
	}

	return res

}
//二叉树的公共祖先
func (bst *BinaryTree) BinaryTreeAncestry(nodeA *Map,nodeB *Map) *Map{
	return bst.binarytreeancestry(bst.Root,nodeA,nodeB)
}
//二叉树的公共祖先
func (bst *BinaryTree) binarytreeancestry(root *Map,nodeA *Map,nodeB *Map) *Map{
	//如果到底了
	if root == nil{
		return nil
	}
	//如果搜索到了 节点
	if root.Level == nodeA.Level || root.Level == nodeB.Level{
		return root
	}
	//上面 2 个条件 是 整个递归的结束条件

	left := bst.binarytreeancestry(root.Left,nodeA,nodeB)
	right := bst.binarytreeancestry(root.Right,nodeA,nodeB)
	//如果 左右 都找到了 返上回公共最小祖先
	if left != nil && right  != nil {
		return root
	}
	//往上返回右边的值
	if left != nil{
		return left
	//往上返回左边找到的值
	}else{
		return right
	}

}

//求二叉树的深度
func (bst *BinaryTree) GetDepth() int{
	return bst.getdepth(bst.Root)
}

//求二叉树的深度
func (bst *BinaryTree) getdepth(root *Map) int{
	if root.Left == nil || root.Right == nil{
		return 1
	}
	leftlen := bst.getdepth(root.Left)
	rightlen := bst.getdepth(root.Left)
	if leftlen > rightlen{
		return leftlen + 1
	}else{
		return rightlen + 1
	}

}

func main(){


	bst := NewBinaryTree()
	arr :=[16]int{74,97,110,101,0,0,0,155,59}
	for i:=0;i< len(arr);i++{
		bst.Add(arr[i])
	}
	//bst.PreOrder()
	//bst.InOrder()
	//bst.PostOrder()
	//fmt.Println("最大值",bst.FindMax().Level)
	//fmt.Println("最小值",bst.FindMin().Level)
	//fmt.Println(bst.String())
	//bst.RemoveMax()
	//bst.RemoveMin()
	//fmt.Println("最大值",bst.FindMax().Level)
	//fmt.Println("最小值",bst.FindMin().Level)
	//bst.Remove(45)
	//bst.InOrder()
	//bst.PostOrder()
	nodea := &Map{27,nil,nil}
	nodeb := &Map{60,nil,nil}
	fmt.Println(bst.StringNoRecursion())
	fmt.Println(bst.BinaryTreeAncestry(nodea,nodeb))
	fmt.Println(bst.GetDepth())


}