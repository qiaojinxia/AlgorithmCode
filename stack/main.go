package main

import "fmt"

/**
 * Created by @CaomaoBoy on 2020/2/22.
 *  email:<115882934@qq.com>
 */

func main(){
	op,err := NewOperator("I+a")
	if err != nil{
		fmt.Println(err)
	}
	value,err := op.Execute([]string{"I","12","a","18"})
	fmt.Println(value)

}