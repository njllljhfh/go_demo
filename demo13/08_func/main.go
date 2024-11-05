package main

import "fmt"

// 闭包

/*
	全局变量特点：
		1. 常驻内存
		2. 污染全局

	局部变量特点：
		1. 不常驻内存
		2. 不污染全局

	闭包：
		1. 可以让一个变量常驻内存
		2. 可以避免变量污染全局
*/

/*
	闭包（非必要，不要使用，不要滥用）
		1. 闭包是指有权访问另一个函数作用域中的变量的函数
		2. 创建闭包的常见方式：在一个函数内部创建另一个函数，通过该函数访问另一个函数内部的变量
*/

// 全局变量
var a = 12

func test() {
	var a = 3 // 局部变量
	fmt.Println(a)
}

//闭包
func adder1() func() int {
	var i = 10 // 局部变量，常驻内存
	return func() int {
		return i + 1
	}
}

func adder2() func(int) int {
	var i = 10 // 局部变量，常驻内存
	return func(y int) int {
		i += y
		return i
	}
}

func main() {
	test()
	fmt.Println(a)
	fmt.Println("------------------------------------------")

	fn := adder1()
	fmt.Println(fn()) // 11
	fmt.Println(fn()) // 11
	fmt.Println(fn()) // 11
	fmt.Println("------------------------------------------")

	fn2 := adder2()
	fmt.Println(fn2(10)) // 20
	fmt.Println(fn2(10)) // 30
	fmt.Println(fn2(10)) // 40
	fmt.Println("------------------------------------------")

}
