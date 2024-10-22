package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 1.Go中定义字符，字符属于int32类型
	var a = 'a'
	fmt.Printf("值：%v，类型：%T\n", a, a) //值：97，类型：int32

	// 原样输出字符
	fmt.Printf("值：%c，类型：%T\n", a, a) //值：a，类型：int32
	fmt.Println("------------------------------------------")

	// 输出字符串里的字符
	str := "this"
	fmt.Printf("值：%v，原样输出：%c，类型：%T\n", str[2], str[2], str[2]) // 值：105，类型：uint8

	// 一个汉字占用3个字节（utf-8） 一个字母占用1个字节
	str2 := "this" //占用4个字节
	// unsafe.Sizeof()无法查看string类型数据的长度，只能查看变量的大小
	fmt.Println(unsafe.Sizeof(str2)) // 16
	fmt.Println(len(str2))           // 4
	fmt.Println("------------------------------------------")

	str3 := "你好go"
	fmt.Println(len(str3)) // 8
	fmt.Println("------------------------------------------")

	// Go 中汉字使用的是utf-8编码，编码后的值是int32类型
	a2 := '国'
	fmt.Printf("值：%v，类型：%T\n", a2, a2) // 值：22269(Unicode编码10进制)，类型：int32
	fmt.Printf("值：%c，类型：%T\n", a2, a2)
	fmt.Println("------------------------------------------")

	// 循环输出字符串里面的字符
	str4 := "你好 golang"
	for i := 0; i < len(str4); i++ {
		// 按 byte 遍历
		fmt.Printf("值：%v，原样输出：%c，类型：%T\n", str4[i], str4[i], str4[i])
	}
	fmt.Println("------------------------------------------")

	// Go语言中，字符有两种
	// 1.uint8类型，或者叫byte类型，代表ASCII码的一个字符
	// 2.rune类型，代表一个UTF-8字符
	// 当需要处理中文、日文或者其他复合字符时，则需要用到rune类型。rune类型实际是一个int32
	// rune类型用来表示utf8字符，一个rune字符由一个或多个byte组成

	str5 := "你好 golang"
	for _, v := range str5 {
		// 按 rune 遍历
		fmt.Printf("值：%v，原样输出：%c，类型：%T\n", v, v, v)
	}
	fmt.Println("------------------------------------------")

	//修改字符串中的字符
	//要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。
	//无论哪种转换，都会重新分配内存，并复制字节数组。
	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "白萝卜a"
	runeS2 := []rune(s2)
	fmt.Printf("%v, %b, %v\n", runeS2[0], runeS2[0], unsafe.Sizeof(runeS2[0]))
	fmt.Printf("%v, %b, %v\n", runeS2[3], runeS2[3], unsafe.Sizeof(runeS2[3]))
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
	fmt.Println("------------------------------------------")
}
