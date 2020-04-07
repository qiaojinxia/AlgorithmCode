package main

import "fmt"

type avltree struct {
	//定义根节点
	root *node
	//包含节点的数量
	size int
}
//定义一颗二叉树
func NewAvTree() *avltree{
	rt := NewAvlNode()
	return &avltree{
		root:rt,
	}
}
type node struct {
	Key int //值
	left *node //左子树
	right *node //右子树
	height int //当前树的高度 = max{left.height,right.height}
}

func NewAvlNode() *node{
	return &node{
		Key: 0,
		left:  nil,
		right: nil,
		height:0,
	}
}
//插入key
func (avlt *avltree) Insert(key int){
	//第一次插入 将key 插入root
	if avlt.size == 0{
		avlt.root.Key = key
	}else{
		avlt.root = insert(avlt.root,key)
	}
	avlt.size ++
	avlt.root = Adjust(avlt.root)
	//avlt.root = adjust(avlt.root)
}

func insert(root *node,key int) *node{
	//递归的 结束条件 当叶子节点时 插入节点
	if root == nil{
		//在叶子节点下面建立新节点
		nd := NewAvlNode()
		nd.Key = key
		//把叶子结点 返回给上一个节点
		return nd
	}
	//往右边 遍历
	if key > root.Key{
		root.right = insert(root.right,key)
	//往左边 遍历
	}else if key < root.Key{
		root.left = insert(root.left,key)
	}
	//返回节点
	return root
}

func Adjust(root *node) *node{
	if root == nil{
		return nil
	}
	//左边节点的高度
	root.left = Adjust(root.left)
	//右边节点的高度
	root.right = Adjust(root.right)

	var left int
	var right int
	if root.left == nil{
		left = -1
	}else{
		left = root.left.height
	}
	if root.right == nil{
		right = -1
	}else{
		right =root.right.height
	}
	balance := left - right
	//左子树 的高度 比右子树的 高
	if balance == 2{
			if root.left.right == nil ||( root.left.left != nil  && root.left.left.height > root.left.right.height) {
				root = root.RightRotate()
			}else{
				root = root.LRRoate()
			}

	}else if balance == -2{
		if root.right.left == nil || ( root.right.right != nil && root.right.right.height > root.right.left.height){
				root = root.LeftRotate()
			}else{
				root = root.RLRoate()
		}
	}
	//计算 叶子结点返回给到 我的高度 然后通过 max{root.right,root.left} + 1 计算出本节点的高度
	root.height = Max(left,right) + 1
	//返回本节点的高度 告诉给上一层
	return root
}


func (avlnode *node) RightRotate() *node{
	tmp :=avlnode.left
	avlnode.left = tmp.right
	tmp.right = avlnode
	return tmp
}
func (avlnode *node) LeftRotate() *node{
	tmp := avlnode.right
	avlnode.right = tmp.left
	tmp.left = avlnode
	return tmp
}

func (avlnode *node) LRRoate() *node{
	avlnode.left = avlnode.left.LeftRotate()
	return avlnode.RightRotate()
}

func (avlnode *node) RLRoate() *node{
	avlnode.right = avlnode.right.RightRotate()
	return avlnode.LeftRotate()
}
//func adjust(root *node) *node{
//	if root == nil{
//		return nil
//	}
//	lnode :=  adjust(root.left)
//	rnode :=  adjust(root.right)
//	var balancefactor int
//	//左右都为空 为叶子节点
//	if lnode == nil && rnode == nil{
//		return root
//	}
//	if lnode == nil{
//		//如果左边为空 计算右边节点的高度
//		 balancefactor =  0 - rnode.height
//	}else if rnode == nil{
//		 balancefactor = lnode.height - 0
//	}
//	fmt.Println("xxx",balancefactor)
//
//	//如果平衡因子 = +-2 了,进行调整
//	//等于 2 的话 树 向右边倾斜 需要右旋 -2向右边倾斜 需要 左旋
//	if balancefactor == 2{
//		root = root.RightRotate()
//	}else if balancefactor == -2{
//		root =root.LeftRotate()
//	}
//	return root
//}

//找到最小节点的上一个节点
func findMin(avlnode *node) *node{
	if avlnode.left == nil{
		return nil
	}
	min := findMin(avlnode.left)
	if min == nil{
		return avlnode
	}
	return avlnode
}
func delelte(avlnode *node,key int) *node{
	if avlnode == nil{
		return nil
	}
	if  avlnode.Key > key{
		avlnode.left =delelte(avlnode.left,key)
	}else if avlnode.Key < key{
		avlnode.right = delelte(avlnode.right,key)
	}else{
		if avlnode.left != nil && avlnode.right != nil {
			//如果 右边节点已进是页子节点
			if avlnode.right.left == nil {
				//替换成 右边的节点的Key
				avlnode.Key = avlnode.right.Key
				//替换成 右边的节点的右子树
				avlnode.right = avlnode.right.right
				//上面2 步 后 当前节点就变成了右节点
			} else {
				//找到 最小节点的上面一个节点
				nodep := findMin(avlnode.right)
				//最小的一个节点
				minnode := nodep.left
				//处理左边最小节点 可能拥有的右节点
				nodep.left = nodep.left.right
				//替换 成最小一个节点的值
				avlnode.Key = minnode.Key
			}
			//一遍 为 nil 直接返回 另一边
		}else if avlnode.left == nil{
			avlnode = avlnode.right
		}else if avlnode.right == nil{
			avlnode = avlnode.left
		}
	}
	return avlnode
}

//中序遍历
func (avlt *avltree) Delete(key int){
	avlt.root = delelte(avlt.root,key)
	avlt.root = Adjust(avlt.root)
}
//中序遍历
func (avlt *avltree) InorderTraversal(){
	inordertraversal(avlt.root)

}
func inordertraversal(root *node){
	if root == nil{
		return
	}
	inordertraversal(root.left)
	fmt.Println(root.Key)
	inordertraversal(root.right)
}

func (avlt *avltree) Find(key int) *node{
	return find(avlt.root,key)
}
func find(avlnode *node,key int) *node{
	if avlnode == nil{
		return nil
	}
	if avlnode.Key > key{
		ret := find(avlnode.left,key)
		return ret
	}else if avlnode.Key < key{
		ret :=find(avlnode.right,key)
		return ret
	}else{
		return avlnode
	}
}
 type ListNode struct {
		Val int
	    Next *ListNode
	}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	mergeNode := &ListNode{
		Val:  0,
		Next: nil,
	}
	res := mergeNode
	for l1 != nil || l2 != nil{
		if l2 != nil && (l1 ==  nil  || l1.Val > l2.Val) {
			mergeNode.Next = l2
			l2 = updateList(l2)

		}else if l1.Next != nil && (l2 ==  nil || l1.Val <= l2.Val){
			mergeNode.Next = l1
			l1 = updateList(l1)
		}
		mergeNode = updateList(mergeNode)
	}
	strx :=  byte('A')
	a := make([]byte,0)
	a = append(a,strx)
	 if cap(a) ==  4{

	 }

	return res.Next
}
func appendRecoder(str byte,sortrecoder []byte){
	index := len(sortrecoder) % cap(sortrecoder)
	sortrecoder = sortrecoder[:0]
	sortrecoder[index] = str
}
func updateList(node *ListNode) *ListNode{
	node = node.Next
	return node
}

func checkInclusion(s1 string, s2 string) bool {
	start := 0
	end := len(s1) -1
	for i:=0;i<len(s2);i++{
		if s2[i] == s1[start] && end == len(s1)-1{
			start += 1
		}else{
			start = 0
		}
		if s2[i] == s1[end] && start == 0{
			end -= 1
		}else{
			end = len(s1)-1
		}
		if start == len(s1)   ||  end == -1{
			return true
		}
	}

	return false

}
func main(){
	fmt.Println(checkInclusion("abc","bbbca"))

	//n1 := &ListNode{Val:  1,Next: nil}
	//n2 := &ListNode{Val:  3,Next: nil}
	//n3 := &ListNode{Val:  4,Next: nil}
	//n4 := &ListNode{Val:  6,Next: nil}
	//n5 := &ListNode{Val:  8,Next: nil}
	//n7 := &ListNode{Val:  10,Next: nil}
	//n1.Next = n2
	//n2.Next = n3
	//n3.Next = n4
	//n4.Next = n5
	//n5.Next = n7
	//n21 := &ListNode{Val:  3,Next: nil}
	//n22 := &ListNode{Val:  13,Next: nil}
	//n23 := &ListNode{Val:  65,Next: nil}
	//n24 := &ListNode{Val:  12,Next: nil}
	//n25 := &ListNode{Val:  43,Next: nil}
	//n26 := &ListNode{Val:  76,Next: nil}
	//n21.Next = n22
	//n22.Next = n23
	//n23.Next = n24
	//n24.Next = n25
	//n25.Next = n26
	//
	//mergeTwoLists(n1,n21)
	//
	//
	//
	a := NewAvTree()
	//i := 3
	//fmt.Println(i/2)
	//a.Insert(50)
	//
	//a.Insert(66)
	//a.Insert(40)
	//a.Insert(21)
	//a.Insert(60)
	//a.Insert(86)
	//a.Insert(55)
	//a.Insert(12)
	//a.Insert(70)
	//a.Insert(31)
	//a.Insert(80)
	//fmt.Println(a.Find(12))
	////a.InorderTraversal()
	//a.Delete(55)
	//fmt.Println(a.Find(12))
	//a.InorderTraversal()


}
/**
 * Created by @CaomaoBoy on 2020/3/24.
 *  email:<115882934@qq.com>
 */
