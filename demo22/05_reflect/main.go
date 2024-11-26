package main

import (
	"fmt"
	"reflect"
)

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(120) // panic: reflect: reflect.Value.SetInt using unaddressable value
	}
}

func reflectSetValue2(x interface{}) {
	// *x = 120 // invalid operation: cannot indirect x (variable of type interface{})
	// - - -

	// v, _ := x.(*int)
	// *v = 120  // panic: runtime error: invalid memory address or nil pointer dereference
	// - - -

	// 通过反射修改值
	v := reflect.ValueOf(x)
	fmt.Println(v.Kind()) // ptr
	fmt.Println(v.Elem().Kind())
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(123)
	} else if v.Elem().Kind() == reflect.String {
		v.Elem().SetString(v.Elem().String() + "666")
	}
}

func main() {
	var a int64 = 100
	// reflectSetValue1(a)

	reflectSetValue2(&a)
	fmt.Println(a) // 123
	fmt.Println("-----------------")

	b := "你好"
	reflectSetValue2(&b)
	fmt.Println(b) // 你好666
}
