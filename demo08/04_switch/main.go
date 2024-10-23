package main

import "fmt"

func main() {
	var extname = ".html"
	switch extname {
	case ".html":
		fmt.Println("text/html")
	case ".css":
		fmt.Println("text/css")
	case ".js":
		fmt.Println("text/javascript")
	default:
		fmt.Println("text/plain")
	}
	fmt.Println("------------------------------------------")

	switch extname2 := ".html"; extname2 {
	case ".html":
		fmt.Println("text/html")
	case ".css":
		fmt.Println("text/css")
	case ".js":
		fmt.Println("text/javascript")
	default:
		fmt.Println("text/plain")
	}
	// fmt.Println(extname2) // 这里会报错（undefined: extname2），extname2变量的作用域在switch语句中
	fmt.Println("------------------------------------------")

	var n = 5
	switch n {
	case 1, 3, 5, 7, 9: // 一个分支可以有多个值
		fmt.Println("奇数")
		// break // 可以不写break，因为switch语句会自动退出
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	}
	fmt.Println("------------------------------------------")

	var age = 30
	switch { // 当分支是条件表达式时，switch 后面不写变量
	case age < 24:
		fmt.Println("好好学习")
	case age >= 24:
		fmt.Println("好好工作")
	case age >= 60:
		fmt.Println("好好享受")
	default:
		fmt.Println("输出错误")
	}
	fmt.Println("------------------------------------------")

	// switch 的穿透 fallthrough
	age = 30
	switch { // 当分支是条件表达式时，switch 后面不写变量
	case age < 24:
		fmt.Println("好好学习")
	case age >= 24 && age < 60:
		fmt.Println("好好工作")
		fallthrough // 穿透
	case age >= 60:
		fmt.Println("注意身体")
		fallthrough
	default:
		fmt.Println("输出错误")
	}
	fmt.Println("------------------------------------------")
}
