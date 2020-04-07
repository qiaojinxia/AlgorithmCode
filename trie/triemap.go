package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

/**
 * Created by @CaomaoBoy on 2020/3/17.
 *  email:<115882934@qq.com>
 */
type trieIndex struct {
	subTrieNodeKey byte      //子节点的索引 字节
	arrindex      int 	   //子节点所在索引
}
//新建 trie 索引结构
func NewtrieIndex() *trieIndex{
	return &trieIndex{
		subTrieNodeKey: 0,
		arrindex:      0,
	}
}
//获取 索引信息
func(ti *trieIndex) getTrieKey() byte{
	return ti.subTrieNodeKey
}
type trieNode struct {
	count		 uint32		  //单词出现次数
	isWord       bool         //是否单词结尾
	subTrieNode  []*trieNode  //所有子节点 list
	subTrieIndex []*trieIndex //记录subTrie排序后的索引
}

//插入 子节点
func(td *trieNode) inserSubTrie(key byte,trienode *trieNode){
	//长度不相等 报错
	if len(td.subTrieNode) != len(td.subTrieIndex){
		panic("error")
	}
	//插入 索引和内容
	td.subTrieNode = append(td.subTrieNode,trienode)
	//初始化一索引信息节点
	nti := NewtrieIndex()
	//索引为 长度 -1
	nti.arrindex = len(td.subTrieNode) -1
	//添加 节点的key
	nti.subTrieNodeKey = key
	//插入
	td.subTrieIndex = append(td.subTrieIndex, nti)

	MergeSort(td.subTrieIndex)

}
////自动排序
//func(trie *Trie) SortItSelf(){
//	trie.sortInterval ++
//	//每插入一定次数 进行一次排序
//	if trie.sortInterval > SortIntervalTimes{
//		trie.sortInterval = 0
//		//对索引 进行 希尔排序
//		ShellSort(trie.root.subTrieIndex)
//	}
//}

//递归进行希尔排序
func sortIndexs(root *trieNode){
	ShellSort(root.subTrieIndex)
	for i:=0;i<len(root.subTrieIndex);i++{
		_,next:=root.getByIndex(i)
		sortIndexs(next)
	}
}

//通过索引 获取 子节点 并调整 索引
func(td *trieNode) delTrieNode(key byte) *trieNode{
	//查找指定Key 对应的index结构
	index := binSearch(td.subTrieIndex,key)
	//没有找到
	if index == -1{
		return nil
	}
	//找到索引 获取 对应索引
	sti := td.subTrieIndex[index].arrindex
	tmpres := td.subTrieNode[sti]
	//通过索引找到结构
	td.subTrieNode = append(td.subTrieNode[:sti],td.subTrieNode[sti+1:]...)
	td.subTrieIndex = append(td.subTrieIndex[:index],td.subTrieIndex[index+1:]...)
	//调整后面节点 索引
	for i:=0;i<len(td.subTrieNode);i++{
		if  td.subTrieIndex[i].arrindex >= sti{
			td.subTrieIndex[i].arrindex = td.subTrieIndex[i].arrindex -1
		}
	}
	return  tmpres
}


//通过索引 获取 子节点
func(td *trieNode) getByIndex(index int) (byte,*trieNode){

	sti := td.subTrieIndex[index].arrindex
	return td.subTrieIndex[index].subTrieNodeKey,td.subTrieNode[sti]
}

//通过二分查找
func binSearch(arr []*trieIndex,data byte) int{
	left := 0
	right := len(arr) -1//最下面最上面
	//针对 数组长度为0的情况
	if right == 0{
		if arr[0].subTrieNodeKey == data{
			return 0
		}else{
			return -1
		}
	}
	for left <= right{
		leftv := float64(data) -float64( arr[left].subTrieNodeKey)
		//如果小于最小值则找不到
		if leftv < 0{
			return -1
		}
		allv  := float64(arr[right].subTrieNodeKey) - float64(arr[left].subTrieNodeKey)
		if leftv > allv{
			return -1
		}
		diff  := float64(right - left)
		mid   := int(float64(left) + leftv /allv * diff)
		if mid < 0 || mid >= len(arr)  {
			return -1
		}
		if  arr[mid].subTrieNodeKey > data{
			right = mid - 1
		}else if arr[mid].subTrieNodeKey < data{
			left = mid + 1
		}else{
			return mid
		}
	}
	return -1
}


//通过 key 获取 节点
func(td *trieNode) getByKey(key byte) *trieNode{
	//查找指定Key 对应的index结构
	index := binSearch(td.subTrieIndex,key)
	//没有找到
	if index == -1{
		return nil
	}
	//找到索引 获取 对应索引
	sti := td.subTrieIndex[index].arrindex
	//通过索引找到结构
	return td.subTrieNode[sti]
}

func MergeSort(arr []*trieIndex){
	if len(arr) <=1{
		return
	}
	//计算位置
	j := len(arr) -1
	for (j > 0 && arr[j-1].getTrieKey() > arr[j].getTrieKey() ){
		swap(arr,j,j-1)
		j --
	}
}

//交换位置
func swap(arr []*trieIndex,i,j int){
	arr[i],arr[j] =arr[j],arr[i]
}
//对内容进行堆排序
func ShellSort(arr []*trieIndex){
	if len(arr) <=1{
		return
	}else{
		//排序 gap
		gap := len(arr) / 2
		for gap >0{
			//计算gap间距的数组
			for i:=0;i<gap;i++{
				ShellSortStep(arr,i,gap)
			}
			//缩小间距
			gap /= 2
		}
	}
}

func ShellSortStep(arr []*trieIndex,start,gap int){
	len := len(arr)
	//按间隔 遍历 所有数组
	for i:=start + gap;i<len;i+=gap{
		//记录 前一个 数组
		j := i - gap
		//由于后面的 会被 前面的插入 所以备份下
		bak := arr[i]
		//如果 前一个数组小于arr 就不听向前插入
		for j >=0  && arr[j].getTrieKey() > bak.getTrieKey()   {
			//如果当前的 比后面的 大 就往后移
			arr[j+gap] = arr[j]
			j -= gap
		}
		//上面的循环完成后 把 bak 放到后面的位置
		arr[j + gap] = bak
	}
}

type Trie struct{
	sortInterval int     //排序 每间隔次
	size int		 //节点大小
	root *trieNode  //根节点
}


//初始化 Trie树
func NewTrie() *Trie{
	node := trieNode{
		isWord:         false,
		count:0,
		subTrieNode:    make([]*trieNode,0),
		subTrieIndex: make([]*trieIndex,0),
	}
	a := Trie{
		size: 0,
		root: &node,
	}
	return &a
}

//如果存在插入
func(trie *Trie) InsertByte(content []byte){
		trie.Insert(content)
}

//如果存在插入
func(trie *Trie) InsertStr(content string){
	trie.Insert([]byte(content))
}

//如果存在插入
func(trie *Trie) InsertInt(content int){
	x := uint32(content)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	trie.Insert(bytesBuffer.Bytes())
}
//往 节点内 添加内容
func(trie *Trie) Insert(content []byte){
	if len(content) ==0{
		panic("error!")
	}
	if  trie.root == nil{
		node := trieNode{
			isWord:         false,
			count:0,
			subTrieNode:    make([]*trieNode,0),
			subTrieIndex: make([]*trieIndex,0),
		}
		trie.root = &node
	}
	var cur *trieNode
	trie.size ++
	for i,v := range content{
		var tmpnode *trieNode
		//如果cur为 nil
		if cur == nil{
			//取出 索引对应的索引
			if trie.root.getByKey(v) == nil || len(trie.root.subTrieNode) ==0{

				tmpnode = &trieNode{
					isWord:         false,
					count:0,
					subTrieNode:    make([]*trieNode,0),
					subTrieIndex: make([]*trieIndex,0),
				}
				//如果是字符串 结尾
				if len(content) -1 == i{
					tmpnode.isWord = true
					tmpnode.count ++
				}
				//更新成当前节点
				cur = tmpnode
				//插入节点
				trie.root.inserSubTrie(v,tmpnode)
			}else{
				cur = trie.root.getByKey(v)
			}

		//如果 cur 已经被初始化
		}else {
			tnode := cur.getByKey(v)
			if tnode == nil{
				tmpnode = &trieNode{
					isWord:      false,
					count:0,
					subTrieNode:    make([]*trieNode,0),
					subTrieIndex: make([]*trieIndex,0),
				}
				//如果是字符串 结尾
				if len(content) -1 == i{
					tmpnode.isWord = true
					tmpnode.count ++
				}
				cur.inserSubTrie(v,tmpnode)
				cur = tmpnode
			}else{
				cur = tnode
				//如果是字符串 结尾
				if len(content) -1 == i{
					cur.count ++
					//重复计次
					cur.isWord = true
				}

			}
		}
	}
}

func (trie *Trie) ContainsStr(content string) bool{
	return trie.contains([]byte(content))
}

//是否包含 str
func (trie *Trie) contains(content []byte) bool{
	if trie.root == nil{
		panic("error nil")
	}
	var cur *trieNode
	for i,v := range content{
		if cur == nil{
			//第一次 从root的next子节点获取
			cur = trie.root.getByKey(v)
		}else{
			//当前节点 的next子节点里获取
			cur = cur.getByKey(v)
		}
		//如果 没有 获取到 认为 没有
		if cur == nil{
			return false
		}
		//如果已经 到达了最后一个节点 但是这个节点没有被标记成 单词
		if len(content) -1 ==  i && cur.isWord == false{
			return false
		}

	}
	return true
}

//遍历查找所有节点内容
func(trie *Trie) PrintAllBytes(deep int,count int){
	if trie.root == nil{
		panic("root nil")
	}
	//res := make([][]byte,0)
	//deepSearch(trie.root,0,&res)
	//for _,v := range res{
	//	fmt.Println(string(v))
	//}
	res := make([][]byte,0)
	deepSearchx(trie.root,deep,nil,&res)
	for _,v := range res {
		if count !=0 {
			ct := BytesToInt(v[len(v)-4:])
			if ct == count{
				fmt.Println("Key:",v[:len(v) -4],"计数:",ct,string(v[:len(v) -4]))
			}

		}else{
			fmt.Println("Key:",v[:len(v) -4],"计数:",BytesToInt(v[len(v)-4:]))
		}

	}

}
//整形转换成字节
func IntToBytes(n uint32) []byte {
	x := uint32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}
//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x uint32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}

//遍历查找所有节点内容
func(trie *Trie) PrintAllStr(deep int){
	if trie.root == nil{
		panic("root nil")
	}
	//res := make([][]byte,0)
	//deepSearch(trie.root,0,&res)
	//for _,v := range res{
	//	fmt.Println(string(v))
	//}
	res := make([][]byte,0)
	deepSearchx(trie.root,deep,nil,&res)
	for _,v := range res {
		fmt.Println("Key:",string(v[:len(v) -4]),"计数:",BytesToInt(v[len(v)-4:]))
	}
}


//深度优先遍历
func deepSearchx(root *trieNode,deep int,res []byte,buf *[][]byte) {
	for _,v := range res {
		if v ==0{
			panic("error")
		}
	}
	if root.isWord{
		rr := make([]byte,0)
		rr = append(rr, res...)
		rr = append(rr, IntToBytes(root.count)...)
		//添加 保存
		*buf = append(*buf,rr)
	}
	for i:=0;i<len(root.subTrieNode);i++{
		//搜索深度
		if deep == 0 {
			return
		}
		key,next:=root.getByIndex(i)
		tmp := make([]byte,0)
		tmp = append(tmp,res...)
		tmp =append(tmp,key)
		deepSearchx(next, deep -1, tmp,buf)
	}
}
//用来输出显示
func deepSearchforDelete(root *trieNode,k byte,res []byte) {
	if len(root.subTrieNode) == 0{
		fmt.Println("clear:",string(k) + string(res))
	}
	for i:=0;i<len(root.subTrieNode);i++{
		key,next:=root.getByIndex(i)
		tmp := make([]byte,0,0)
		tmp = append(tmp,res...)
		tmp =append(tmp,key)
		deepSearchforDelete(next,k, tmp)
	}
}

//是否包含 str
func (trie *Trie) Delete(content string) bool{
	if trie.root == nil{
		panic("error nil")
	}
	var cur *trieNode
	for i,v := range []byte(content){
		if cur == nil{
			//第一次 从root的next子节点获取
			cur = trie.root.getByKey(v)
		}else{
			//当前节点 的next子节点里获取
			cur = cur.getByKey(v)
		}
		//如果 没有 获取到 认为 没有
		if cur == nil{
			return false
		}
		//如果已经 到达了最后一个节点 但是这个节点没有被标记成 单词
		if len(content) -1 ==  i && cur.isWord == false{
			return false
		}

	}
	//计数器 1次 就修改为false
	if cur.count == 1{
		cur.isWord = false//找到了 就删除
	}else{
		cur.count --
	}

	trie.Adjust()//删除没用节点
	return true
}


//将被删除的节点清除掉
func(trie *Trie) Adjust(){
	adjust(trie.root,trie.root,0,0)
}
//将被删除的节点清除掉
func adjust(root *trieNode,ptr *trieNode,key byte,deep int){

	//如果 没有next了 并且 这个节点 没有 记录单词
	if len(root.subTrieNode) == 0 && root.isWord == false{
		deepSearchforDelete(ptr.getByKey(key),key,nil)
		//删除子节点
		ptr.delTrieNode(key)
		return
	}else if len(root.subTrieNode) == 0{
		return
	}
	//如果 是 单词记录节点
	for i:=0;i< len(root.subTrieNode);i++{
		k,subnode := root.getByIndex(i)
		if deep == 0{
			key = k
		}
		//记录下这个节点
		if root.isWord {
			ptr = root //记录父节点
			key = k//记录子节点 key
		}
		adjust(subnode,ptr,key,deep + 1)
	}

}

//匹配 字符串
func(trie *Trie) Match(content string){


}

//自动提示
func(trie *Trie) Suggess(content string){
	if trie.root == nil{
		panic("error nil")
	}
	var cur *trieNode
	savebyte := make([]byte,0,0)
	for i,v := range []byte(content){
		//保存前面输入的字节
		savebyte = append(savebyte, v)
		if cur == nil{
			//第一次 从root的next子节点获取
			cur = trie.root.getByKey(v)
		}else{
			//当前节点 的next子节点里获取
			cur = cur.getByKey(v)

		}
		//如果 没有 获取到 认为 没有
		if cur == nil{
			fmt.Println("no srarch")
			return
		}
		//如果已经 到达了最后一个节点 但是这个节点没有被标记成 单词
		if len(content) -1 ==  i  && cur!= nil{
			res := make([][]byte,0)
			deepSearchx(cur,-1,nil,&res)
			for _,v := range res {
				new := savebyte
				new = append(new,v[:len(v) -4] ...)
				//fmt.Println("提示:",string(new),"计数:",BytesToInt(v[len(v)-4:]))
			}
			fmt.Println("查询到:",len(res))
		}

	}
}
func main(){
	t := NewTrie()
	t.InsertStr("mriemap")
	t.InsertStr("mriema")
	t.InsertStr("tEST00001")
	t.InsertStr("tEST0000")
	t.InsertStr("tEST0001")
	t.InsertStr("zEST0")
	//t.InsertStr("caomao")
	//t.InsertStr("caomao")
	//t.InsertStr("SADASDASDSADASD")
	////t.Insert("apple");
	//t.PrintAllStr(-1)
	//t.InsertStr("app");
	//fmt.Println(t.ContainsStr("caomao"))
	//t.Delete("caomao")
	//t.Delete("caomao")
	//fmt.Println("--------")
	//fmt.Println(t.ContainsStr("caomao"))
	//t.PrintAllStr(-1)
	//t.Suggess("tE")
	//sortarr := make([]int,0)
	//rand.Seed(time.Now().UnixNano())
	//for i := 0; i < 10000000; i++  {
	//	sortarr = append(sortarr, rand.Intn(100000))
	//}
	//now := time.Now()
	////for i:=0;i<len(sortarr);i++{
	////	t.InsertInt(sortarr[i])
	////}
	//fmt.Println(time.Now().Sub(now).Milliseconds(),"毫秒")

	now := time.Now()
	file ,err:= os.Open("/Users/qiao/go/src/qqsort/trie/CSDNpass.txt")
	if err != nil{
		panic(err)
	}
	rf := bufio.NewReader(file)
	i := 0
	for{
		line,_,err := rf.ReadLine()
		if err == io.EOF{
			break
		}
		for _,m := range strings.Split(string(line)," "){
			val := strings.TrimSpace(m)
			i+=1
			if len(val) == 0{
				continue
			}
			t.InsertStr(val)
		}

	}


	//t.Suggess([]byte{1})
	fmt.Println(time.Now().Sub(now).Milliseconds(),"毫秒")

	fmt.Println("key个数",t.size)
	for{
		fmt.Println("请输入一个字符串:")
		//读键盘
		reader := bufio.NewReader(os.Stdin)
		//以换行符结束
		str, _ := reader.ReadString('\n')
		fmt.Println("当前Trie树单词数:",t.size)
		nowx := time.Now()
		//t.Suggess(strings.TrimSpace(str))
		fmt.Println("是否包含:",t.ContainsStr(str))
		fmt.Println(time.Now().Sub(nowx).Microseconds(),"us")
	}
	time.Sleep(time.Second * 30)

}