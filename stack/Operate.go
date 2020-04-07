package main

import (
	"math"
	"strconv"
	"strings"
)

/**
 * Created by @CaomaoBoy on 2020/2/22.
 *  email:<115882934@qq.com>
 */

const(
	AVAIABLE_CODE = "+-*/^√"
	AVAIABLE_DECIMAL_CODE = "1234567890.E"
	//参数
	AVAIABLE_PARAMETER_CODE ="abcdefghijklmnopqrstuvwxyz"
)
func IsAvaiablecode(c string) bool{
	return strings.IndexAny(AVAIABLE_CODE,c) != -1
}

func isBelongToDecimal(c string)bool{
	return strings.IndexAny(AVAIABLE_DECIMAL_CODE,c) != -1
}

func isParmertercode(c string)bool{
	return strings.IndexAny(AVAIABLE_PARAMETER_CODE,c) != -1
}

func exesingleExpression(left float64,right float64,exp string)float64{
	if exp == "+"{
		return left + right
	}else if exp == "-"{
		return left - right
	}else if exp == "*"{
		return left * right
	}else if exp == "/"{
		return left / right
	}else if exp == "^"{
		return math.Pow(left,right)
	}else if exp == "√"{
		return math.Pow(left,1/right)
	}
	return 0.0
}
func changeParameter(Parameter string,str[]string)string{
	for i:=0;i<len(str);i+=2{
		if Parameter == str[i]{
			return str[i+1]
		}
	}
	return Parameter
}
type Operator struct {
	sentence string // 1 + (2 * 3) 文字表达式
	opers []string	//表达式存储 2 * 3 1 +6
	suffixExpression []string //后缀表达式
}
//新建四则运算类
func NewOperator(sentence string)(*Operator,error){
	a:= &Operator{
		sentence:         sentence,
		opers:            nil,
		suffixExpression: nil,
	}
	a.opers = make([]string,0)
	return a,nil
}
//初始化
func (this *Operator) init(){
	return
}
//中缀表达式转化为后缀表达式,自动具备了顺序
func (this *Operator)setShuffixExpression()error{
	return nil
}
func (this *Operator)Execute(str []string)(value float64,err error){
	temp := NewStack()
	for i:=0;i<len(this.suffixExpression);i++{
		st := changeParameter(this.suffixExpression[i],str)
		if val,err := strconv.ParseFloat(strings.TrimSpace(st),64);err != nil{
			temp.Push(val)
		}else{
			exp := this.suffixExpression[i]
			if exp =="I"{
				v1 := temp.Pop() //取得弹出的数据
				temp.Push(M2int64(v1))
			}else {
				rights := temp.Pop()
				right :=M2float64(rights)//数据转化

				lefts := temp.Pop()
				left :=M2float64(lefts)//数据转化
				temp.Push(exesingleExpression(left,right,exp))//递归调用

			}
		}
	}
	value = M2float64(temp.Pop())
	return
}
