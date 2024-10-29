package main

import "fmt"

func main() {
	// 声明切片
	var slice1 []int
	fmt.Printf("%v - %T - 长度:%v\n", slice1, slice1, len(slice1))

	var slice2 = []int{1, 2, 34, 45}
	fmt.Printf("%v - %T - 长度:%v\n", slice2, slice2, len(slice2))

	var slice3 = []int{1: 2, 2: 4, 5: 6}
	fmt.Printf("%v - %T - 长度:%v\n", slice3, slice3, len(slice3))
	fmt.Println("------------------------------------------")

	var slice4 []int
	fmt.Println(slice4)
	fmt.Println(slice4 == nil) // true: go中声明切片以后，切片的默认值是nil

	var slice5 = []int{1, 2, 34, 45}
	fmt.Println(slice5 == nil) // flase
	fmt.Println("------------------------------------------")

	// 切片的遍历
	var strSlice = []string{"北京", "上海", "广州", "深圳", "成都", "重庆"}
	for i := 0; i < len(strSlice); i++ {
		fmt.Println(strSlice[i])
	}
	fmt.Println("=====")

	for k, v := range strSlice {
		fmt.Println(k, v)
	}
	fmt.Println("------------------------------------------")

	// 基于数组定义切片
	a := [5]int{55, 56, 57, 58, 59}
	fmt.Printf("%v - %T\n", a, a)

	b := a[:] // 获取数组里的所有值
	fmt.Printf("%v - %T\n", b, b)

	c := a[1:4]
	fmt.Printf("%v - %T\n", c, c) // [56 57 58] - []int

	d := a[2:]
	fmt.Printf("%v - %T\n", d, d) // [57 58 59] - []int

	e := a[:3]
	fmt.Printf("%v - %T\n", e, e) // [55 56 57] - []int
	fmt.Println("------------------------------------------")

	//基于切片定义切片
	a2 := []string{"北京", "上海", "广州", "深圳", "成都", "重庆"}

	b2 := a2[1:]                    // 获取数组里的所有值
	fmt.Printf("%v - %T\n", b2, b2) // [上海 广州 深圳 成都 重庆] - []string
	fmt.Println("------------------------------------------")

	// 切片的长度和容量
	// 切片的长度 就是它所包含的元素个数
	// 切片的容量 是从它的第一个元素开始数，到其底层数组最后一个元素的个数
	//
	// s 的底层数组是 [2, 3, 5, 7, 11, 13]
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("长度%d 容量%d\n", len(s), cap(s)) // 长度6 容量6

	s2 := s[2:]                                            // [5 7 11 13]
	fmt.Printf("s2=%v, 长度%d 容量%d\n", s2, len(s2), cap(s2)) // 长度4 容量4

	s3 := s[1:3]                                           //[3 5]
	fmt.Printf("s3=%v, 长度%d 容量%d\n", s3, len(s3), cap(s3)) // 长度2 容量5

	s4 := s[:3]                                            //[2 3 5]
	fmt.Printf("s4=%v, 长度%d 容量%d\n", s4, len(s4), cap(s4)) // 长度3 容量6

	// s5 := s4[1:7] // 报错
	s5 := s4[1:len(s)]
	fmt.Println(s5) // [3 5 7 11 13]
	fmt.Println("------------------------------------------")

	// 使用make函数构造切片
	var sliceA = make([]int, 4, 8)
	fmt.Printf("s4=%v, 长度%d 容量%d\n", sliceA, len(sliceA), cap(sliceA))
	fmt.Println("------------------------------------------")
}
