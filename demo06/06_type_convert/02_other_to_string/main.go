package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 1. 通过 fmt.Sprintf() 将其他类型转换成 string 类型
	// 注意：Sprintf 使用中需要注意转换的格式，int为%d，float为%f，bool为%t, byte为%c

	var i int = 20
	var f float32 = 12.456
	var t bool = true
	var b byte = 'a'

	str1 := fmt.Sprintf("%d", i)
	fmt.Printf("值：%v, 类型：%T\n", str1, str1)

	str2 := fmt.Sprintf("%f", f)
	fmt.Printf("值：%v, 类型：%T\n", str2, str2)

	str3 := fmt.Sprintf("%t", t)
	fmt.Printf("值：%v, 类型：%T\n", str3, str3)

	str4 := fmt.Sprintf("%c", b)
	fmt.Printf("值：%v, 类型：%T\n", str4, str4)
	fmt.Println("------------------------------------------")

	// 2.通过 strconv 将其他类型转换成 string 类型
	var i2 int = 20
	str5 := strconv.FormatInt(int64(i2), 10) // 10 表示十进制
	fmt.Printf("值：%v, 类型：%T\n", str5, str5)

	/*
		参数1：要转换的值
		参数2：格式化类型('f' - 带小数的浮点数，例如 123.456)
		参数3：小数点后保留的位数
		参数4：格式化的类型 传入 64 表示 float64，传入 32 表示 float32
	*/
	var f2 float32 = 20.231313
	str6 := strconv.FormatFloat(float64(f2), 'f', 4, 64)
	fmt.Printf("值：%v, 类型：%T\n", str6, str6)

	var t2 bool = true
	str7 := strconv.FormatBool(t2) // 不常用
	fmt.Printf("值：%v, 类型：%T\n", str7, str7)

	var b2 byte = 'a'
	str8 := strconv.FormatUint(uint64(b2), 10) // 不常用
	fmt.Printf("值：%v, 类型：%T\n", str8, str8)
}
