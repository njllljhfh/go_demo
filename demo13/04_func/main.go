package main

import "fmt"

// 自定义类型
type calc func(int, int) int // 表示定义一个calc的类型

type myInt int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func main() {
	var c calc
	c = sub
	fmt.Printf("c的类型：%T\n", c) // c的类型：main.calc
	fmt.Println(c(10, 5))

	f := sub                   // 没有对f显示指定类型
	fmt.Printf("f的类型：%T\n", f) // f的类型：func(int, int) int

	fmt.Println(f(15, 5))
	fmt.Println("------------------------------------------")

	var a int = 10
	var b myInt = 20
	fmt.Printf("a的类型：%T, b的类型：%T\n", a, b)
	fmt.Println(a + int(b)) // 类型转化
	fmt.Println("------------------------------------------")
}
