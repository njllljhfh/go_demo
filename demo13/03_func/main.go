package main

import "fmt"

var a = "全局变量"

func run() {
	var b = "局部变量"
	fmt.Printf("run方法 a=%v\n", a)
	fmt.Printf("run方法 b=%v\n", b)
}

func main() {
	fmt.Println(a)

	run()
}
