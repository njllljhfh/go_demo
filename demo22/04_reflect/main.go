package main

import (
	"fmt"
	"reflect"
)

/*
通过反射，获取变量的原始值
*/

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	kind := v.Kind() // 类型种类(底层的类型)，实际用的比较多
	switch kind {
	case reflect.Int64:
		fmt.Printf("int64 的原始值:%v\n", v.Int())
	case reflect.Float32:
		fmt.Printf("float32 的原始值:%v\n", v.Float())
	case reflect.Float64:
		fmt.Printf("float64 的原始值:%v\n", v.Float())
	case reflect.String:
		fmt.Printf("string 的原始值:%v\n", v.String())
	default:
		fmt.Printf("还没有判断这个类型")
	}
}

func main() {
	var a float32 = 3.14
	var b int64 = 100
	var c string = "你好xxx"
	reflectValue(a)
	reflectValue(b)
	reflectValue(c)
}
