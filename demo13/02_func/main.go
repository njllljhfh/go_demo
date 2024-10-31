package main

import (
	"fmt"
	"sort"
)

// 升序
func sortIntAsc(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	return slice
}

// 降序
func sortIntDesc(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] < slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	return slice
}

func main() {
	slice := []int{12, 34, 37, 35, 556, 36, 2}
	sortIntAsc(slice)
	fmt.Println(slice)

	fmt.Println(sortIntDesc(slice))
	fmt.Println("------------------------------------------")

	/*
		当对 Go 中的切片（slice）进行排序时，切片底层数组的内容会发生变化。
		排序操作直接修改了切片底层数组中的元素顺序，
		因此切片的所有引用都会反映出新的排序结果。
	*/
	numbers := []int{5, 3, 4, 1, 2}
	// 创建一个指向同一底层数组的切片
	refSlice := numbers[1:4]
	// 对原切片进行排序
	sort.Ints(refSlice)
	fmt.Println("Sorted numbers:", numbers)   // 输出: [5 1 3 4 2]
	fmt.Println("Reference slice:", refSlice) // 输出: [1 3 4]

}
