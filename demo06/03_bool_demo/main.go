package main

import "fmt"

func main() {
	/*
		1.布尔类型变量的默认值是false。
		2.Go 语言中不允许将整型强制转换为布尔型。
		3.布尔型无法参与数值运算，也无法与其他类型进行转换。
	*/

	var flag bool = true
	fmt.Printf("%v--%T\n", flag, flag)
	fmt.Println("------------------------------------------")

	// 布尔类型变量的默认值是false
	var b bool
	fmt.Printf("%v\n", b)

	// string类型变量的默认值是""
	var s string
	fmt.Printf("%v\n", s)

	// int类型变量的默认值是0
	var i int
	fmt.Printf("%v\n", i)

	// float类型变量的默认值是0
	var f float32
	fmt.Printf("%v\n", f)
	fmt.Println("------------------------------------------")

	// 在Go语言中，不允许将整型强制转换为布尔型。(错误示例)
	// var a = 1
	// if a {
	// 	fmt.Println("true")
	// }
}
