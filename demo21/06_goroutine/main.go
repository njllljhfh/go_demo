package main

import (
	"fmt"
	"sync"
	"time"
)

/*
需求：要统计 1-120000 的数字中哪些是素数。
素数：只能被1和它本身整除的数（1不是素数）

协程实现
协程1：统计 1-30000
协程2：统计 30001-60000
协程3：统计 60001-90000
协程4：统计 90001-120000
*/
var wg sync.WaitGroup

// test 统计素数
func test(start, end int) {
	defer wg.Done()
	for num := start; num <= end; num++ {
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
}

func main() {
	total := 120003
	goroutineNum := 4
	// 每个协程最少处理的数据个数
	per := total / goroutineNum
	fmt.Println(per)

	start := time.Now().UnixMilli()

	for i := 1; i <= goroutineNum; i++ {
		start := (i-1)*per + 1
		if start == 1 {
			start = 2
		}
		end := i * per
		if total-end < per {
			end = total
		}
		fmt.Printf("协程%v: start=%v, end=%v\n", i, start, end)
		wg.Add(1)
		go test(start, end)
	}
	wg.Wait()

	end := time.Now().UnixMilli()
	fmt.Printf("协程执行的时间为: %v 毫秒\n", end-start)
}
