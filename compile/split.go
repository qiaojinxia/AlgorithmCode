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
	const example = `a1 b2 c3
d e3 f
h1 i4 j`
	var (
		column int
		row int
		scan scanner.Scanner
		strArray [3][3]string
	)
	scan.Init(strings.NewReader(example)) //初始化
	scan.Whitespace ^= 1 << '\t' | 1 <<'\n' | 1 << ' ' //切割字符
	var token rune
	for token = scan.Scan();token!= scanner.EOF;token=scan.Scan(){
		switch token {
		case '\n':
			row ++
			column = 0
		case ' ':
			column++
		default:
			strArray[row][column] = scan.TokenText()//扫描字符串

		}
	}
	fmt.Println(strArray)

}