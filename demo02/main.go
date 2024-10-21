package main

import "fmt"

func main() {
	// fmt.Println("你好 golang")  // 会换行
	// fmt.Print("你好 golang")
	// fmt.Printf("你好 golang")

	// fmt.Print("A", "B", "C")   // ABC
	// fmt.Println("A", "B", "C") // A B C

	// var a = "aaa" //go语言中，变量定义了以后必须要使用
	// fmt.Println(a)
	// fmt.Printf("%v", a) // 格式化输出

	// var a int = 10
	// var b int = 3
	// var c int = 5
	// fmt.Println("a=", a, "b=", b, "c=", c)
	// fmt.Printf("a=%v b=%v c=%v", a, b, c)

	a := 10
	b := 20
	c := 30
	fmt.Println("a=", a, "b=", b, "c=", c)
	fmt.Printf("a=%v b=%v c=%v\n", a, b, c)
	fmt.Printf("a的类型是：%T\n", a)
}
