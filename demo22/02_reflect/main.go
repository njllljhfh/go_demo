package main

import (
	"fmt"
	"reflect"
)

type myInt int
type Person struct {
	Name string
	Age  int
}

// 反射获取任意变量的类型
func reflectFn(x interface{}) {
	v := reflect.TypeOf(x)
	// v.Name() // 类型名称
	// v.Kind() // 类型种类(底层的类型)，实际用的比较多
	fmt.Printf("类型:%v, 类型名称:%v, 类型种类:%v\n", v, v.Name(), v.Kind())
}

func main() {
	a := 10
	b := 23.4
	c := true
	d := "你好"
	reflectFn(a) // 类型:int, 类型名称:int, 类型种类:int
	reflectFn(b) // 类型:float64, 类型名称:float64, 类型种类:float64
	reflectFn(c) // 类型:bool, 类型名称:bool, 类型种类:bool
	reflectFn(d) // 类型:string, 类型名称:string, 类型种类:string

	fmt.Println("-----------------------")

	var e myInt = 34
	var f = Person{
		Name: "张三",
		Age:  20,
	}
	reflectFn(e) // 类型:main.myInt, 类型名称:myInt, 类型种类:int
	reflectFn(f) // 类型:main.Person, 类型名称:Person, 类型种类:struct
	fmt.Println("-----------------------")

	// 数组、切片、Map、指针 的类型名称是空的
	var h = 25
	reflectFn(&h) // 类型:*int, 类型名称:, 类型种类:ptr

	var i = [3]int{1, 2, 3}
	reflectFn(i) // 类型:[3]int, 类型名称:, 类型种类:array

	var j = []int{11, 22, 33}
	reflectFn(j) // 类型:[]int, 类型名称:, 类型种类:slice

	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	reflectFn(m)
}
