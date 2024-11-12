package main

import (
	"fmt"
	"itying/calc"
	T "itying/tools" // 给包起别名
)

// import _ "itying/tools" // 匿名导入包，导入包后不使用包中的任何函数和变量

// go mod init 项目根目录名称 (生成go.mod文件)

/*
	包名为main的包为应用程序的入口包，编译后会生成可执行文件，而编译不包含 main 包的源代码文件则不会生成可执行文件
	一个项目里只能有一个main包，main包中只能有一个main函数

	在运行时，最后被导入的包会最先被初始化并调用其 init 函数
	如 a 包导入了 b 包，b 包导入了 c 包，那么 c 包的 init 函数会最先被调用，然后 b 包的 init 函数会被调用，最后 a 包的 init 函数会被调用
*/

func main() {
	sum := calc.Add(10, 2)
	fmt.Println(sum)

	fmt.Printf("Age = %d\n", calc.Age)

	res := T.Mul(10, 2)
	fmt.Println(res)

	T.PringInfo()

	arr := []int{5, 4, 3, 2, 1}
	T.Sort(arr)
	fmt.Println(arr)
}

// init 函数，在main函数执行前执行，被系统自动调用
func init() {
	fmt.Println("main init...")
}
