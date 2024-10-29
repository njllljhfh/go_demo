package main

import (
	"fmt"
)

func main() {
	// 切片 是 引用数据类型
	slice1 := []int{1, 2, 3, 45}
	slice2 := slice1 // 切片的赋值 是 引用传递
	slice2[0] = 11
	fmt.Println(slice1)
	fmt.Println(slice2)
	/*
		[11 2 3 45]
		[11 2 3 45]
	*/
	fmt.Println("------------------------------------------")

	// copy 方法 拷贝切片
	slice3 := []int{1, 2, 3, 45}
	slice4 := make([]int, 4, 4)
	copy(slice4, slice3)
	slice4[0] = 111
	fmt.Println(slice3)
	fmt.Println(slice4)
	/*
		[1 2 3 45]
		[111 2 3 45]
	*/
	fmt.Println("------------------------------------------")

	// 删除索引为 2 的元素
	a := []int{30, 31, 32, 33, 34, 35, 36, 37}
	fmt.Printf("a=%v, 长度%d 容量%d\n", a, len(a), cap(a))

	a = append(a[:2], a[3:]...)
	fmt.Printf("a=%v, 长度%d 容量%d\n", a, len(a), cap(a))
	fmt.Println("------------------------------------------")

	// 改变字符串中的字符
	s1 := "big"
	byteStr := []byte(s1)
	fmt.Println(byteStr)
	byteStr[0] = 'p'
	fmt.Println(string(byteStr))
	fmt.Println("===")

	s2 := "你好golang"
	runeStr := []rune(s2)
	fmt.Println(runeStr)
	runeStr[0] = '大'
	runeStr[len(runeStr)-1] = 'G'
	fmt.Println(string(runeStr))
}
