package main

import "fmt"

func main() {
	// 基本数据类型 和 数组 都是值类型
	var arr1 = [...]int{1, 2, 3}
	arr2 := arr1
	arr2[0] = 100
	fmt.Println(arr1) // arr1 没有变化
	fmt.Println(arr2)
	fmt.Println("------------------------------------------")

	// 切片，是引用类型
	var slice1 = []int{1, 2, 3}
	slice2 := slice1
	slice2[0] = 200
	fmt.Println(slice1) // slice1 发生了变化
	fmt.Println(slice2)
	fmt.Println("------------------------------------------")
}
