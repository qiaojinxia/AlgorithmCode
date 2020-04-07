package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

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
	file,_ = parser.ParseFile(fileset,"code.go",example,parser.ImportsOnly)
	var spec *ast.ImportSpec
	for _,decalrations := range file.Decls{
		pos := decalrations.Pos()
		relativepos := fileset.Position(pos)
		abpos := fileset.PositionFor(pos,false)

		var gokeyword string
		gokeyword = "func"
		if gen,ok :=decalrations.(*ast.GenDecl);ok{
			gokeyword = gen.Tok.String()
		}
		fmt.Println(gokeyword)
	}
}