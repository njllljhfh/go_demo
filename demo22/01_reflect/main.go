package main

import (
	"fmt"
	"reflect"
)

/*
反射是指在程序运行期间对程序本身进行访问和修改的能力。
正常情况下，程序在编译时，变量被转换为内存地址，变量名不会被编译器写入到可执行部分。在运行程序时，程序无法获取自身的信息。
支持反射的语言可以在程序编译期将变量的反射信息，如字段名称、类型信息、结构体信息等整合到可执行文件中，并给程序提供接口访问反射信息，
这样就可以在程序运行期获取类型的反射信息，并且有能力修改它们。

1.反射可以在程序运行期间动态获取变量的各种信息，比如变量的类型，类别
2.如果是结构体，通过反射还可以获取结构体本身的信息，比如结构体的字段、结构体的方法。
3.通过反射，可以修改变量的值，可以调用关联的方法。

Go语言中的变量是分为两部分的：
    - 类型信息：预先定义好的元信息。
    - 值信息：程序运行过程中可动态变化的。
在Go语言的反射机制中，任何接口值都是由一个 具体类型 和 具体类型的值 两部分组成的。

reflect包提供的两个重要的函数：
reflect.TypeOf
reflect.ValueOf
*/

type myInt int
type Person struct {
	Name string
	Age  int
}

// 反射获取任意变量的类型
func reflectFn(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Println(v)
}

func main() {
	a := 10
	b := 23.4
	c := true
	d := "你好"
	reflectFn(a) // int
	reflectFn(b) // float64
	reflectFn(c) // bool
	reflectFn(d) // string
	fmt.Println("-----------------------")

	var e myInt = 34
	var f = Person{
		Name: "张三",
		Age:  20,
	}
	reflectFn(e) // main.myInt
	reflectFn(f) // main.Person
	fmt.Println("-----------------------")

	var h = 25
	reflectFn(&h) // *int
}
