package main

import "fmt"

/**
 * Created by @CaomaoBoy on 2020/2/17.
 *  email:<115882934@qq.com>
 */
func HeapSort(arr []int,topn int) []int{
	length := len(arr)
	for i:=1;i<=length;i++{
		lastlen := length -i
		HeapSortMax(arr,lastlen)
		if i< length {
			arr[0],arr[lastlen] = arr[lastlen],arr[0]
		}
		if topn - i ==0{
			break
		}
	}
	return arr[len(arr) - topn:]
}

//每次从堆中 取一个最大数 放到 i处


func HeapSortMax(arr[] int,length int) []int{
	if length <= 1{
		return arr
	}else{
		depth := length/2 -1
		for i:=depth;i>=0;i--{
			topmax := i //最大的在i位置
			leftchild := 2 * i + 1
			rightchild := 2 * i + 2
			if leftchild <= length  && arr[leftchild] > arr[topmax]{
				topmax = leftchild
			}
			if rightchild <= length  && arr[rightchild] >arr[topmax]{
				topmax = rightchild
			}
			if topmax != i{
				arr[i],arr[topmax] = arr[topmax],arr[i]
			}

		}
		return arr
	}
}

func main(){
	arr := []int{1,34,54,12,32,12,65,23,67}
	fmt.Println(HeapSortMax(arr,len(arr)))
}