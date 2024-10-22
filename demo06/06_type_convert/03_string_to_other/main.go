package main

import (
	"fmt"
	"strconv"
)

func main() {
	// string类型装换成整型
	str := "123456"
	fmt.Printf("值：%v, 类型：%T\n", str, str)

	/*
		ParseInt()函数的参数：
		1.要转换的字符串
		2.表示字符串的进制，2~36，0、1分别代表二进制和八进制
		3.表示转换的结果的bit大小，32 表示 int32，64 表示 int64

		返回值：
		1.转换后的整型值（转换失败返回 0）
		2.可能的错误
	*/
	num, _ := strconv.ParseInt(str, 10, 64)
	fmt.Printf("值：%v, 类型：%T\n", num, num)
	fmt.Println("------------------------------------------")

	/*
		ParseFloat()函数的参数：
		1.要转换的字符串
		2.表示转换的结果的bit大小，32 表示 float32，64 表示 float64

		返回值：
		1.转换后的浮点型值（转换失败返回 0）
		2.可能的错误
	*/
	str2 := "123.456"
	num2, _ := strconv.ParseFloat(str2, 64)
	fmt.Printf("值：%v, 类型：%T\n", num2, num2)
	fmt.Println("------------------------------------------")

	// 不建议把 string 类型转换为 bool
	b, _ := strconv.ParseBool("true")
	fmt.Printf("值：%v, 类型：%T\n", b, b)

	b, _ = strconv.ParseBool("xxxxxxxxx")
	fmt.Printf("值：%v, 类型：%T\n", b, b)
}
