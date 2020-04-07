package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

/**
 * Created by @CaomaoBoy on 2020/2/22.
 *  email:<115882934@qq.com>
 */

func M2int64(i interface{})int64{
	if i == nil{
		panic(errors.New("类型为空,无法判断!"))
	}
	switch t := i.(type){
	case int64:
		return t
	case int32:
		return int64(t)
	case float32:
		return int64(t)
	case float64:
		return int64(t)
	case byte:
		return int64(t)
	default:
		return Stoi64(M2String(i))
	}

}


func M2float64(i interface{})float64{
	if i == nil{
		panic(errors.New("类型为空,无法判断!"))
	}
	switch t := i.(type){
	case int64:
		return float64(t)
	case int32:
		return float64(t)
	case float32:
		return float64(t)
	case float64:
		return t
	case byte:
		return float64(t)
	default:
		return Stof64(M2String(i))
	}

}
func M2String(i interface{}) string{
	if i == nil{
		panic(errors.New("类型为空,无法判断!"))
	}
	switch t := i.(type){
	case string:
		return t
	default:
		return fmt.Sprintf("%v",i)
	}
}
func Stoi64(v string,def ...int64) int64{
	if n,err := strconv.ParseInt(strings.TrimSpace(v),0,0);err ==nil{
		return n
	}else if len(def) >0{
		return def[0]
	}else{
		panic(errors.New("int64位数据转换"))
	}

}

func Stof64(v string,def ...float64) float64{
	if n,err := strconv.ParseFloat(strings.TrimSpace(v),0);err ==nil{
		return n
	}else if len(def) >0{
		return def[0]
	}else{
		panic(errors.New("float64位数据转换"))
	}

}