package main

import "fmt"

func main() {
	// 1.最简单的if语句
	flag := true
	if flag {
		fmt.Println("flag=true")
	}
	fmt.Println("------------------------------------------")

	// 2.if语句的初始化语句
	if age := 34; age > 20 {
		fmt.Println("成年人")
	} else {
		fmt.Println("未成年人")
	}
	// fmt.Println("成年人", age) // 这里会报错（undefined: age），age变量的作用域在if语句中
	fmt.Println("------------------------------------------")
}
