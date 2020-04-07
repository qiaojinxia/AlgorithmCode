package main

import "fmt"

/**
 * Created by @CaomaoBoy on 2020/2/18.
 *  email:<115882934@qq.com>
 */

func insertsort(arr []int) []int{
	if len(arr) <= 1{
		return arr
	}

			j := 4 -1
			tmp :=arr[4]
			for j >= 0 && tmp < arr[j]{
				arr[j+1] = arr[j]
				j --
				fmt.Println(arr)
			}
			arr[j+1] = tmp



	return arr

}

func main(){
	arr := []int{33,34,54,76,36,32,34,65,23,67}
	fmt.Println(insertsort(arr))
}