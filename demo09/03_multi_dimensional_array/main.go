package main

import "fmt"

func main() {
	var arr = [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(arr)
	fmt.Println(arr[0])
	fmt.Println(arr[0][1])
	fmt.Println("------------------------------------------")

	for _, v1 := range arr {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
	fmt.Println("------------------------------------------")

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Println(arr[i][j])
		}
	}
	fmt.Println("------------------------------------------")

	// 自动推导数组的长度
	var arr2 = [...][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(arr2)
}
