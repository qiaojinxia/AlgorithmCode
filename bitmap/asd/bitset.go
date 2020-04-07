package main

import (
	"fmt"
	"github.com/spaolacci/murmur3"
	"math"
)

/**
 * Created by @CaomaoBoy on 2020/2/24.
 *  email:<115882934@qq.com>
 */
const (
	//这里 ptrBits 可以知道int 在系统中表示的大小byte大小 (^uintptr(0) >> 63) 在 32位 为 0 在 64位 为 1
	//ptrBits = 32 << uint( (^uintptr(0) >> 63)) //^uintptr(0) 表示 64位0 取反  也就是 2 << 63 -1 = 18446744073709551615
	////减1是因为 从0 开始
	//byteModeMask = ptrBits - 1 //能记录的长度 上面 ptrBits 是 64 这里 从0 开始 所以减1
	//byteShift =(1 << 7 + ptrBits) >> 6	//  8位 / 2 << 7 64 位就是 2 << 63
	byteModeMask = 7	//2 ^3 -1
	byteShift = 3 // 2 ^ 3
)
type BitSet interface {
	Get(i int)bool	//提取数据
	Set(i int)		//设置数据
	Unset(i int)	//取消设置
	SetBool(i int,b bool)//设置
}
type Bytes [] byte //指针 可以指向 表示 2 << 63 的大小整数

func NewBytes(numBits int) Bytes{
	// numBits 范围,byteModeMask 起始  如果 numBits = 0    byteModeMask >> byteShift 63 >> 6 = 0
	return make(Bytes,(numBits+ byteModeMask) >> byteShift)
}

//提取数据
func (p Bytes)Get(i int)bool{
	return p[uint(i) >> byteShift] & (1 << (uint(i) & byteModeMask)) != 0

}

//提取数据
func (p Bytes)Set(i int){
	// (uint(i)&byteModeMask  10  % 63 等价于 i(10) & byteModeMask(63)
	//uint(i) >> byteShift 这个等价于 10 / (1 << 63)  计算第几个uintptr
	//|= 一位一位的取或运算
	//所以下面代码 就是对  一个 64位0101的数 对指定为设置为1
	//下面 用|= 和 ^= 是一样的 区别是 如果对同一个数添加2次 用^=2次 就会成为0
	p[uint(i) >> byteShift] |= 1 << (uint(i)&byteModeMask)
}


func (p Bytes)UnSet(i int){
	// 001000 001000 &^= 如果右侧是0，则左侧数保持不变 如果右侧是1，则左侧数一定清零 所以会被清零
	p[uint(i) >> byteShift] &^= 1 << (uint(i)&byteModeMask)
}


func (p Bytes)SetBool(i int,b bool){
	if b{
		p.Set(i)
	}
	p.UnSet(i)
}
//增长数据
func (p * Bytes)Grow (numBits int){
	ptrs := *p
	//等价于 10 + 8  /  8
	targetlen := (numBits + byteModeMask) >> byteShift
	//计算出需要多少个容纳 更搭的数
	missing := targetlen - len(ptrs)
	if missing >0 && missing <= targetlen{
		//添加一个新的 uint64
		*p = append(ptrs,make(Bytes,missing)...)
	}
}
func main(){
	p :=NewBytes(10)
	p.Set(11)
	p.Set(14)
	p.Grow(31)
	fmt.Println(p.Get(11))
	fmt.Println(p.Get(31))
	var  data []byte
	data = []byte{1}
	haser := murmur3.New128() //128的hash 2^128
	haser.Write(data)	//2个整数
	v1,v2 := haser.Sum128()
	fmt.Println(v1,v2)
	fmt.Println(math.Log(2))
}

-1 * n * ln^p/ln2^2