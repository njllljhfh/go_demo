package main

import "fmt"

func main() {
	//整型之间的转换
	var a int8 = 20
	var b int16 = 40
	fmt.Println(int16(a) + b)
	fmt.Println("------------------------------------------")

	//浮点型之间的转换
	var a2 float32 = 20
	var b2 float64 = 40
	fmt.Println(float64(a2) + b2)
	// 转换的时候，建议从低精度向高精度转换，否则可能会有精度损失
	fmt.Println("------------------------------------------")

	//浮点型和整型之间的转换
	var a3 float32 = 20.23
	var b3 int = 40
	fmt.Println(a3 + float32(b3))
	// 转换的时候，建议从int向float转换，否则可能会有精度损失
	fmt.Println("------------------------------------------")
}
