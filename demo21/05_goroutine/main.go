package main

import (
	"fmt"
	"time"
)

/*
	需求：要统计 1-120000 的数字中哪些是素数。
	用for循环实现
*/

func main() {
	// for循环实现
	start := time.Now().UnixMilli()
	for num := 2; num <= 120003; num++ {
		flag := true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Printf("%v 是素数\n", num)
		}
	}
	end := time.Now().UnixMilli()
	fmt.Printf("for循环执行的时间为: %v 毫秒\n", end-start)
}
