package main

import (
	"fmt"
	"sort"
)

func main() {
	// map 排序
	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90
	map1[11] = 100
	map1[2] = 13
	fmt.Println(map1)

	for key, val := range map1 {
		fmt.Println(key, val)
	}
	fmt.Println("------------------------------------------")

	// 按照 kye 升序 输出
	var keySlice []int
	for key, _ := range map1 {
		keySlice = append(keySlice, key)
	}
	fmt.Println(keySlice)
	sort.Ints(keySlice)
	fmt.Println(keySlice)

	for _, k := range keySlice {
		fmt.Printf("k=%v, v=%v\n", k, map1[k])
	}
}
