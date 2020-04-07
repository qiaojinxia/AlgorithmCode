package main

import (
	"errors"
	"fmt"
)

/**
 * Created by @CaomaoBoy on 2020/2/24.
 *  email:<115882934@qq.com>
 */

type BitMap struct {
	size int //大小
	vals []byte //字节
}
func NewBitMap(size int)  BitMap{
	//位置从 1 开始
	length := size / 8 + 1
	return BitMap {
		size: size,
		vals: make([]byte,length,length),
	}
}
func (b BitMap) Size() int{
	return b.size
}

func (b BitMap)CheckRange(i int)error{
	if i> b.Size(){
		return errors.New("index too big")
	}
	if i < 1{
		return errors.New("index too small")
	}
	return nil
}

func (b BitMap)toggle(i int){
	p := i >> 3
	remainder := i - (p * 8)
	fmt.Println("XXXX")
	if remainder == 1{
		b.vals[p] = b.vals[p]^1

	}else{
		b.vals[p] = b.vals[p]^(1 << uint(remainder))
	}
}

func (b BitMap)Set(i int)error{
	if x := b.CheckRange(i);x != nil{
		return x
	}
	val,err := b.IsSet(i)
	if err != nil{
		return err
	}
	if val{
		return nil
	}
	b.toggle(i)
	return nil
}

func (b BitMap)UnSet(i int)error{
	val,err := b.IsSet(i)
	if err != nil{
		return err
	}
	if val{
		b.toggle(i)
	}
	return nil
}
func (b BitMap)IsSet(i int) (bool,error){
	if x := b.CheckRange(i);x != nil{
		return false,x
	}
	p := i >> 3 // /8
	remainder := i - (p * 8)
	if remainder == 1{
		return b.vals[p] > b.vals[p]^1,nil
	}
	return b.vals[p] > b.vals[p] ^(1 << uint((remainder))),nil

}
func (b BitMap)Values() ([]int,error){
	list := make([]int,0,b.Size())
	for i:=1;i<b.Size();i++ {
		val, err := b.IsSet(i)
		if err != nil {
			return nil, err
		}
		if val {
			list = append(list, i)
		}
	}
	return list,nil
}

func main(){
	b := NewBitMap(100)
	b.Set(8)
	b.Set(12)
	b.Set(43)
	b.Set(23)
	b.Set(65)
	b.Set(12)
	b.Set(33)
	vals,_:= b.Values()
	fmt.Println(vals)
	c := 32
	d := 63
	fmt.Println(c &d)



}