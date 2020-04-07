package main

import "fmt"

/**
 * Created by @CaomaoBoy on 2020/2/24.
 *  email:<115882934@qq.com>
 */


func setbit(bm []byte,i int){
	i  -= 1
	p := i / 8
	l := i % 8
	bytem := bm[p] //这就取出8位了
	fmt.Printf("改变前状态:%08b 第 %d 位要置为1 \n", bytem,l+ 1)
	bytem = bytem ^ 1 << l
	fmt.Printf("改变后    :%08b\n", bytem)
	bm[p] = bytem
}

func getbit(bm []byte,i int) bool{
	i  -= 1
	p := i / 8
	l := i % 8
	bytem := bm[p] //这就取出8位了
	return bytem > bytem ^ 1 << l

}

func iter(bm []byte) []int{
	list := make([]int,0,len(bm))
	for i:=1;i<=len(bm) * 8;i++ {
		val := getbit(bm,i)
		if val {
			list = append(list, i)
		}
	}
	return list

}
func UnSet(bm []byte,i int){
	fmt.Println("删除数字:",i)
	val := getbit(bm,i)
	if val {
		setbit(bm,i)
	}
}

func main(){

	bitsmap := make([]byte,2,2)
	setbit(bitsmap,8)
	setbit(bitsmap,16)
	fmt.Println(getbit(bitsmap,8))
	fmt.Println(getbit(bitsmap,2))
	fmt.Println(iter(bitsmap))
	UnSet(bitsmap,8)
	fmt.Println(iter(bitsmap))
	a := 32 << uint(uintptr(0) >> 63)
	fmt.Println(a)
	n := (^uintptr(0))
	fmt.Println(n)
}