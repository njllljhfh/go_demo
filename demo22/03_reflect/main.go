package main

import (
	"fmt"
	"reflect"
)

// 反射获取任意变量的类型
func reflectValue(x interface{}) {
	// var num = 10 + x // invalid operation: 10 + x (mismatched types int and interface{})

	// 类型断言
	b, _ := x.(int)
	var num = 10 + b
	fmt.Println(num)

	// 反射
	// v := reflect.ValueOf(x)
	// var n = v + 12 // cannot convert 12 (untyped int constant) to type struct{typ_ *abi.Type; ptr unsafe.Pointer; reflect.flag}
	// fmt.Println(n)
	v := reflect.ValueOf(x)
	var m = v.Int() + 12 // 反射获取变量的原始值
	fmt.Println(m)
}

func main() {
	var a = 13
	reflectValue(a)
}
