package main

import "fmt"

// 自定义类型
type myInt int

// 类型别名
type myFloat = float64

func main() {
	var a myInt = 10
	fmt.Printf("a=%v, a的类型：%T\n", a, a) // a=10, a的类型：main.myInt

	var b myFloat = 10
	fmt.Printf("b=%v, a的类型：%T\n", b, b) // b=10, a的类型：float64
}
