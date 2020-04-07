package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

/**
 * Created by @CaomaoBoy on 2020/2/25.
 *  email:<115882934@qq.com>
 */


func main(){
	const example =`
	//要扫描的代码
	if a > 10{
	fmt.Println("你妹")	
	}
	`
	fmt.Println(example)
	var sI scanner.Scanner //扫描器
	sI.Init(strings.NewReader(example)) //扫描器
	sI.Filename ="code"
	var tok rune
	for tok = sI.Scan();tok!= scanner.EOF;tok = sI.Scan(){
		fmt.Printf("%s:%s \n",sI.Position,sI.TokenText())
	}


}