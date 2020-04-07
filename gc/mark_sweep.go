package main

import (
	"fmt"
)

/**
 * Created by @CaomaoBoy on 2020/2/18.
 *  email:<115882934@qq.com>
 */

type Heap struct {
	root *Obj //根节点
	size int
	objs []*Obj
}

type ObjHead struct {
	myname string
	ismark bool
	ref []*Obj
}
//是否标记
func ( oh *ObjHead) ifmark() bool{
	return oh.ismark
}

//是否标记
func (oh *ObjHead) help(){
	fmt.Printf("%s 被标记了 快去救救它吧!\n ",oh.myname)
}
//添加引用
func ( oh *ObjHead) addRef(obj *Obj){
		fmt.Printf("%s  在山里 捡到一个 %s ! !\n ",oh.myname,obj.oh.myname)
	if oh.ref == nil{
		oh.ref = make([]*Obj,0,1)
	}
	oh.ref = append(oh.ref, obj)

}

//添加引用
func ( oh *ObjHead) subRef(obj *Obj){
	for i,m := range oh.ref {
		if m ==  obj{
			fmt.Printf(" %s 抛弃了 %s !\n",oh.myname,obj.oh.myname)
			oh.ref[i] = nil
			return
		}
	}
	fmt.Printf(" %s 和 %s 没有引用关系 不需要解除!\n",oh.myname,obj.oh.myname)
}

//对象
type Obj struct {
	//对象实体
	data interface{}
	//对象头
	oh ObjHead
}

//生成对象
func NewObj(name string) *Obj{
	head :=ObjHead{myname:name,ismark:false,ref:nil}
	fmt.Printf("伟大的 %s 诞生了....\n",name)
	return &Obj{
		data: nil,
		oh:   head,
	}
}

//清理分为2 标记 和 清除
func(h *Heap) Sweep(){
	//标记
	mark(h.root)
	//清理
	sweep(h.objs)
}


//创建堆
func NewHeap() *Heap{
	return &Heap{
		size:0,
		objs:make([]*Obj,0,10)}
}

//创建对象
func(h *Heap)NewObj(name string) *Obj{
	obj := NewObj(name)
	usenilspace := false
	//利用已经被回收的内存
	for i,space := range h.objs{
		if space == nil{
			h.objs[i] = obj
			usenilspace = true
		}
	}
	if usenilspace == false{
		h.objs = append(h.objs,obj)
	}
	h.size ++
	return obj
}

//从根节点 扫描所有引用对象
func  mark(obj *Obj) {
	if obj == nil{
		return
	}
	//从root开始扫描所有引用对象
	if (obj.oh.ismark == false){
		obj.oh.ismark = true
	}
	for _,i:= range obj.oh.ref  {
		mark(i)
	}
}

//回收内存
func  sweep(objs []*Obj) {
	for i,m := range objs{
		if m.oh.ismark == false{
			m.oh.help()
			//回收对象 指针被修改为nil 遍历数组时可以复用
			objs[i] = nil
		}else{
			//修改为false 为下次回收做准备
			m.oh.ismark = false
		}
	}
}

func main(){
	hp := NewHeap()
	root := NewObj("爷爷")
	//设置更节点
	hp.root = root

	o1 := hp.NewObj("大娃")
	o2 := hp.NewObj("二娃")
	o3 := hp.NewObj("三娃")
	o4 := hp.NewObj("四娃")
	o5 := hp.NewObj("五娃")
	o6 := hp.NewObj("六娃")
	hp.NewObj("七娃")//七娃 生成了对象没被引用

	root.oh.addRef(o1) //相当于 root = o1
	o1.oh.addRef(o2)
	o1.oh.addRef(o3)
	o1.oh.addRef(o4)
	root.oh.addRef(o5)
	o2.oh.addRef(o6)
	root.oh.subRef(o1) //大娃 关联了 二娃 三娃 四娃 二娃 关联了 六娃 二三四六娃 会随着大娃的解出引用同时解除引用
	hp.Sweep()


}
