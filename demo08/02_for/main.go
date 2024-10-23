package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}
	// fmt.Println(i)  // i 在 for 循环外不可见
	fmt.Println("------------------------------------------")

	i := 1
	for ; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Printf("i = %d\n", i)
	fmt.Println("------------------------------------------")

	i2 := 1
	for i2 <= 10 {
		fmt.Println(i2)
		i2++
	}
	fmt.Println("------------------------------------------")

	// 无限循环
	i3 := 1
	for {
		if i3 <= 10 {
			fmt.Println(i3)
		} else {
			break
		}
		i3++
	}
}
