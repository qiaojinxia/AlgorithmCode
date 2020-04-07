package main

import (
	"fmt"
	"math"
	"time"
)

/**
 * Created by @CaomaoBoy on 2020/4/1.
 *  email:<115882934@qq.com>
 */

func findRepeatNumber(nums []int) int {
	//arr := make([]int,len(nums),len(nums))
	for i,val := range nums {
		if nums[i] != i{
			if nums[val] == val{
				return val
			}else{
				nums[i],nums[val] = nums[val],val
			}
		}
	}
	return -1
}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	length := len(matrix)
	middle := float64(length / 2)
	point := int(middle)
	for point >= 0 && point <= length -1{
			ls := matrix[point]
			ptr :=len(ls) /2
			statue := 0
			for ptr >= 0 && ptr < len(ls) {
					if ls[ptr] > target  && (statue ==0 || statue ==1){
						statue = 1
						ptr -= 1
					}else if ls[ptr] < target && (statue ==0 || statue ==2){
						statue = 2
						ptr += 1
					}else if ls[ptr] ==target{
						return true
					}else{
						break
					}
			}

			if float32(middle)/2 < 0.5{
				break
			}
			middle = middle /2.0
			if statue == 1{
				point += int(math.Ceil(middle))
			}else if statue == 2 {
				point -= int(math.Ceil(middle))
			}
	}
	return false

}
func findNumberIn2DArrayx(matrix [][]int, target int) bool {
	if len(matrix) == 0{
		return false
	}
	ptr := len(matrix[0])-1
	for _,v := range matrix{
		for i:=ptr;i>=0;i--{
			if target > v[i]{
				ptr = i
				break
			}else if target < v[i]{
				continue
			}else{
				return true
			}
		}
	}
	return false
}


func main(){
	//a := []int{5,6,10,14}
	//b := []int{6,10,13,18}
	//c := []int{10,13,18,19}

	a := []int{1,4,7,11,15}
	b := []int{2,5,8,12,19}
	c := []int{3,6,9,16,22}
	d := []int{10,13,14,17,24}
	e:= []int{18,21,23,26,30}
	xx := make([][]int,0,0)
	xx = append(xx, a)
	xx = append(xx, b)
	xx = append(xx, c)
	xx = append(xx, d)
	xx = append(xx, e)



	//xx := make([][]int,0,0)
	//a :=[]int{5}
	//b :=[]int{6}
	//xx = append(xx, a)
	//xx = append(xx, b)




	before := time.Now()
	fmt.Println(findNumberIn2DArray(xx,14))
	fmt.Println("spend times:",time.Now().Sub(before))


	//nums :=[]int{0, 1, 2, 3, 4, 11, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	//fmt.Println(findRepeatNumber(nums))
}