package main

import "fmt"

/**
 * Created by @CaomaoBoy on 2020/2/17.
 *  email:<115882934@qq.com>
 */

 func selectMax(arr[] int) int{
 	length := len(arr)
 	if length <=1{
 		return arr[0]
	}else{
		max := arr[0] //选取第一个
		for i:=1;i<length;i++{
			if arr[i] > max{
				max = arr[i]//选取商户组中最大值
			}
		}
		return max
	}
 }


 //选择排序每次找到数组最小的 然后把其他数据进行交换
 func SelectorSort(arr[]int) []int{
 	length := len(arr)
 	if length <= 1 {
 		return arr
	}else{
		for i:=0;i<length-1;i++{
			min := i
			for j:= i+1;j<length;j++{
				if arr[min] > arr[j]{
					min = j//存储最小的
				}
			}
			if i != min{
				arr[i],arr[min] = arr[min],arr[i] //交换
			}
		}
	}
	return arr
 }


 func main(){
 	arr := []int{1,34,54,12,32,12,65,23,67}
 	fmt.Println(SelectorSort(arr))
 }