package main

import "fmt"

/**
 * Created by @CaomaoBoy on 2020/2/17.
 *  email:<115882934@qq.com>
 */

func GetMax(arr[] int)int{
	for j:=1;j<len(arr);j++{
		if arr[j-1] < arr[j]{
			arr[j-1],arr[j] = arr[j],arr[j-1]
		}
		fmt.Println(arr)
	}
	return arr[len(arr)-1]
}


//适用场景: 对已经有序的数组进行插入
func BubbleSort(arr[] int) []int{
	for i:=0;i<len(arr);i++{
		for j:=i+1;j<len(arr);j++{
			if arr[i] > arr[j]{
			arr[i],arr[j] = arr[j],arr[i]
			}
		}
		fmt.Printf("\n第%d次排序, %v",i,arr)
	}
	return arr
}

func main() {
	arr := []int{1, 34, 54, 12, 32, 12, 65, 23, 67,11}
	fmt.Println("\nres:",BubbleSort(arr))
}
