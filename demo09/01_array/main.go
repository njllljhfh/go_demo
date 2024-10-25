package main

import "fmt"

func main() {
	// 数组的长度也是数组类型的一部分，因此，不同长度的数组为不同的类型
	var arr1 [3]int
	var arr2 [4]int
	var strArr [3]string
	fmt.Printf("arr1:%T, arr2:%T, strArr:%T\n", arr1, arr2, strArr)
	fmt.Println("------------------------------------------")

	// 数组初始化-方法1
	var arr3 [3]int
	fmt.Println(arr3)
	arr3[0] = 23
	arr3[1] = 10
	arr3[2] = 24
	fmt.Println(arr3)
	fmt.Println("------------------------------------------")

	// 数组初始化-方法2
	var arr4 = [3]int{1, 2, 3}
	fmt.Println(arr4)
	fmt.Println("------------------------------------------")

	// 数组初始化-方法3
	var arr5 = [...]int{1, 2, 3, 4, 5}
	arr5[0] = 100
	fmt.Println(arr5)
	fmt.Println(len(arr5))
	fmt.Println("------------------------------------------")

	// 数组初始化-方法4
	arr6 := [...]int{0: 1, 1: 10, 2: 20, 5: 50}
	fmt.Println(len(arr6))
	fmt.Println(arr6)
	fmt.Println("------------------------------------------")

	// 数组的循环遍历
	arr7 := [...]string{"php", "nodejs", "golang", "java"}
	for i := 0; i < len(arr7); i++ {
		fmt.Println(arr7[i])
	}
	fmt.Println("------------------------------------------")

	for k, v := range arr7 {
		fmt.Printf("key:%v, value:%v\n", k, v)
	}
	fmt.Println("------------------------------------------")

	// 从数组[1, 3, 5, 7, 8] 中找到和为8的两个元素的下标分别是（0,3）和（1,2）。
	arr8 := [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(arr8); i++ {
		for j := i + 1; j < len(arr8); j++ {
			if arr8[i]+arr8[j] == 8 {
				fmt.Printf("(%v,%v)\n", i, j)
			}
		}
	}
}
