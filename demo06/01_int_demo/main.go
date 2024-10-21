package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var num int = 10
	fmt.Printf("num=%v, 类型：%T\n", num, num)

	// var num2 int8 = 130  // 130 超出了 int8 的范围
	var num2 int16 = 130 // 130 超出了 int8 的范围
	fmt.Printf("num2=%v\n", num2)
	fmt.Println("------------------------------------------")

	// unsafe.Sizeof() 查看占用多少 字节
	var num3 int8 = 15
	fmt.Printf("num3=%v, 类型:%T, 大小:%d 字节\n", num3, num3, unsafe.Sizeof(num3))

	var num4 uint8 = 10
	fmt.Printf("num4=%v\n", num4)
	fmt.Println("------------------------------------------")

	// int 类型转换
	var a1 int32 = 10
	var a2 int64 = 21
	fmt.Println(int64(a1) + a2) // 把a1转换为64位
	fmt.Println(a1 + int32(a2)) // 把a1转换为64位

	var n1 int16 = 130
	fmt.Println(int8(n1)) // -126
	fmt.Println("------------------------------------------")

	num5 := 30
	fmt.Printf("num5=%v, 类型:%T\n", num5, num5) // num5 为 int类型:在64位的计算机上表示的就是int64，在32位的计算机上表示的就是int32
	fmt.Println(unsafe.Sizeof(num5))           // 8
	fmt.Println("------------------------------------------")

	num6 := 18
	fmt.Printf("num6=%v\n", num6) //%v 原样输出
	fmt.Printf("num6=%d\n", num6) //%d 十进制输出
	fmt.Printf("num6=%b\n", num6) //%d 二进制输出
	fmt.Printf("num6=%o\n", num6) //%d 八进制输出
	fmt.Printf("num6=%x\n", num6) //%d 十六进制输出
}
