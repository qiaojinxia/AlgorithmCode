package main

import (
	"fmt"
	"runtime"
	"sync"
)

/**
 * Created by @CaomaoBoy on 2020/2/17.
 *  email:<115882934@qq.com>
 */

func shellsort(arr []int) []int{
	if len(arr)  <= 1 {
		return arr
	}else{
		gap := len(arr)/2
		for gap > 0{
			for i:=0;i<gap;i++{
				ShellSortStep(arr,i,gap)
			}
			gap/=2
		}
		return arr
	}
}

func ShellSortStep(arr[] int,start int,gap int){
	length := len(arr)
	for i:= start + gap;i<length;i+=gap{
		j := i - gap
		bak := arr[i]
		for j >=0 && bak < arr[j] {
			arr[j + gap] = arr[j]
			j -= gap
		}
		arr[j + gap] =bak


	}


}
func gourtinesort(arr []int){
	if len(arr) <2 || arr == nil{
		return
	}
	cpunum := runtime.NumCPU()
	wg := sync.WaitGroup{}
	for gap := len(arr);gap >0;gap/=2{
		wg.Add(cpunum)
		ch := make(chan int,1000)
		go func(){
			for k:=0;k<gap;k++{
				ch <- k
			}
			close(ch)
		}()
		for k:=0;k<cpunum;k++{
			go func() {
				for v:= range ch{
					fmt.Println(v,"XXX")
					ShellSortStep(arr,v,gap)
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
	fmt.Println(arr)

}

func main(){
	arr := []int{1,34,54,12,32,12,65,23,67}
	gourtinesort(arr)
}