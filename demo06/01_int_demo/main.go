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
	fmt.Printf("num2 = %v\n", num2)

	// unsafe.Sizeof() 查看占用多少 字节
	var num3 int8 = 15
	fmt.Printf("num3=%v, 类型:%T, 大小:%d 字节\n", num3, num3, unsafe.Sizeof(num3))

}
