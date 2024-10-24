package main

import (
	"fmt"
	"unsafe"

	"github.com/shopspring/decimal"
)

func main() {
	var a float32 = 3.12
	// %f 默认保留小数点后6位，四舍五入
	fmt.Printf("值：%v--%f, 类型%T\n", a, a, a)
	fmt.Println(unsafe.Sizeof(a)) // 占用4个字节
	fmt.Println("------------------------------------------")

	var a2 float64 = 3.12
	// %f 默认保留小数点后6位
	fmt.Printf("值：%v--%f, 类型%T\n", a2, a2, a2)
	fmt.Println(unsafe.Sizeof(a2)) // 占用8个字节
	fmt.Println("------------------------------------------")

	var c float64 = 3.1415926535
	fmt.Printf("%v--%f--%.2f\n", c, c, c)
	fmt.Printf("%.4f\n", c) // 保留4位小数
	fmt.Println("------------------------------------------")

	// 64位的系统中，Go语言中浮点数默认是 float64
	f1 := 3.1415926535
	fmt.Printf("%f--%T\n", f1, f1)
	fmt.Println("------------------------------------------")

	// 科学计数法
	var f2 = 3.14e2 // 表示f2等于 3.14 * 10^2
	fmt.Printf("%v--%T\n", f2, f2)
	var f3 = 3.14e-2 // 表示f2等于 3.14 / 10^2
	fmt.Printf("%v--%T\n", f3, f3)
	fmt.Println("------------------------------------------")

	// float 精度丢失问题
	var f4 float64 = 1129.6
	fmt.Println(f4 * 100) // 112959.99999999999
	fmt.Println("------------------------------------------")

	var num1 float64 = 3.1
	var num2 float64 = 4.2
	fmt.Println(num1 + num2) // 7.300000000000001
	// 解决精度丢失问题
	// 初始化go.mod: go mod init 02_float_demo
	// 安装decimal包: go get github.com/shopspring/decimal
	d1 := decimal.NewFromFloat(num1).Add(decimal.NewFromFloat(num2))
	fmt.Println(d1)
}
