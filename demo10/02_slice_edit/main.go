package main

import "fmt"

func main() {
	var sliceA = make([]int, 4, 8)
	sliceA[0] = 10
	fmt.Println(sliceA)
	fmt.Println("------------------------------------------")

	sliceB := []string{"php", "java", "go"}
	sliceB[2] = "golang"
	fmt.Println(sliceB)
	fmt.Println("------------------------------------------")

	var sliceC []int
	fmt.Printf("sliceC=%v, 长度%d 容量%d\n", sliceC, len(sliceC), cap(sliceC)) // 长度3 容量6
	fmt.Println("------------------------------------------")

	// go 中没法通过下标的方式给切片扩容
	// 切片扩容要用到 append 方法
	var sliceD []int
	sliceD = append(sliceD, 12)
	fmt.Printf("sliceD=%v, 长度%d 容量%d\n", sliceD, len(sliceD), cap(sliceD)) // 长度1 容量1
	sliceD = append(sliceD, 24)
	fmt.Printf("sliceD=%v, 长度%d 容量%d\n", sliceD, len(sliceD), cap(sliceD)) // 长度2 容量2

	// 一次追加多个数据
	sliceD = append(sliceD, 1, 2, 3)
	fmt.Printf("sliceD=%v, 长度%d 容量%d\n", sliceD, len(sliceD), cap(sliceD))
	fmt.Println("------------------------------------------")

	// append 方法 合并切片
	sliceE := []string{"php", "java"}
	sliceF := []string{"nodejs", "go"}
	sliceE = append(sliceE, sliceF...)
	fmt.Println(sliceE)
	fmt.Println("------------------------------------------")

	// append的扩容策略
	var sliceG []int
	for i := 0; i < 10; i++ {
		sliceG = append(sliceG, i)
		fmt.Printf("sliceG=%v, 长度%d 容量%d\n", sliceG, len(sliceG), cap(sliceG))
	}
}
