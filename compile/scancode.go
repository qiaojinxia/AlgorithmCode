package main

import (
	"fmt"
	"go/token"
)
import "go/ast"
import "go/parser"
/**
 * Created by @CaomaoBoy on 2020/2/25.
 *  email:<115882934@qq.com>
 */

func main(){
	const example = `
package main
import(
"fmt"
"strings"
)
func main(){
	fmt.Println(strings.Split("A#B#C","#"))
}
`
	var fileset *token.FileSet
	fileset = token.NewFileSet()
	var file *ast.File //新建文件
	file,_ = parser.ParseFile(fileset,"",example,parser.ImportsOnly)
	var spec *ast.ImportSpec
	for _,spec = range file.Imports{
		fmt.Println(spec.Path.Value) //打印引用的包
		fmt.Println(spec.Path)
	}



}