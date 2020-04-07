package main

import "fmt"

/**
 * Created by @CaomaoBoy on 2020/2/19.
 *  email:<115882934@qq.com>
 */

type Heap struct {
	root *Obj //根节点
	size int
	objs []*Obj
}

//创建堆
func NewHeap() *Heap{
	return &Heap{
		size:0,
		objs:make([]*Obj,0,10)}
}
//创建堆
func (hp *Heap)ScanHeap(){
	for i,m:= range hp.objs{
		if m.oh.ref_cnt == 0{
			fmt.Printf("%s 计数器 为0 已经牺牲了!\n",m.oh.myname)
			hp.objs[i] = nil
		}else{
			fmt.Printf("%s 存活确认!\n",m.oh.myname)
		}
	}

}

type ObjHead struct {
	ref_cnt int
	myname string
	ref []*Obj
}
//对象
type Obj struct {
	//对象实体
	data interface{}
	//对象头
	oh ObjHead
}

func NewObj(name string,isroot bool) *Obj{
	oh := ObjHead{
		ref_cnt: 0,
		myname:  name,
		ref:     nil,
	}
	if isroot == true{
		oh.ref_cnt = 1
	}
	return &Obj{
		data: nil,
		oh:  oh,
	}
}

//设置对象引用
func (o *Obj)addRef(index int,obj *Obj){
	if len(o.oh.ref) == 0{
		o.oh.ref = make([]*Obj,0,10)
		if index != -1{
			fmt.Println("引用对象不存在!")
			return
		}
	}
	if index == -1{
		o.oh.ref = append(o.oh.ref, obj)
		updata_ptr(nil,obj)
	}else{
		if o.oh.ref[index] != nil{
			ptr := o.oh.ref[index]
			updata_ptr(&ptr,obj)
		}

	}
}
//计数器加1
func inc_ref_cnt(obj *Obj){
	if obj == nil{
		return
	}
	fmt.Printf("%s 计数器 加 1\n",obj.oh.myname)
	obj.oh.ref_cnt ++
}

//计数器减1
func dec_ref_cnt(obj *Obj){
	obj.oh.ref_cnt --
	fmt.Printf("%s 计数器 减 1\n",obj.oh.myname)
	//如果 当前技术器 为 0那么就变成 垃圾了 对其 引用的对象 引用减1
	if (obj.oh.ref_cnt == 0 || obj == nil){
		for _,v := range obj.oh.ref{
			dec_ref_cnt(v)
		}
	}
}

func updata_ptr(ptr **Obj,obj *Obj){
	inc_ref_cnt(obj)
	if ptr != nil{
		dec_ref_cnt(*ptr)
		*ptr = obj
	}
}

//创建对象
func(h *Heap)NewObj(name string) *Obj{
	var obj *Obj
	if h.size == 0{
		obj = NewObj(name,true)
	}else{
		obj = NewObj(name,false)
	}
	usenilspace := false
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
func main(){
	hp := NewHeap()
	root := hp.NewObj("爷爷")
	o1 := hp.NewObj("大娃")
	o2 := hp.NewObj("二娃")
	o3 := hp.NewObj("三娃")
	o4 := hp.NewObj("四娃")
	o5 := hp.NewObj("五娃")
	o6 := hp.NewObj("六娃")
	o7 := hp.NewObj("七娃")//七娃 生成了对象没被引用
	root.addRef(-1,o1)
	root.addRef(-1,o2)
	o2.addRef(-1,o3) //添加 二娃 第 0个引用
	o2.addRef(-1,o4)//添加 二娃 第 1个引用
	o4.addRef(-1,o5)
	o4.addRef(-1,o6)
	o2.addRef(1,nil)  //把 二娃引用 四娃置为nil
	//下面就是经典的 循环引用了
	o6.addRef(-1,o7)
	o7.addRef(-1,o6)

	//o5.addRef(0,o6)
	hp.ScanHeap()





}
