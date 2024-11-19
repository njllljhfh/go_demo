package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	协程

	本节内容，实际项目中经常会用到
*/

var wg sync.WaitGroup

func test1() {
	for i := 0; i < 10; i++ {
		fmt.Println("test1...", i)
		time.Sleep(time.Millisecond * 50)
	}
	wg.Done() // 协程计数器-1
}

func test2() {
	for i := 0; i < 10; i++ {
		fmt.Println("test2...", i)
		time.Sleep(time.Millisecond * 50)
	}
	wg.Done() // 协程计数器-1
}

func main() {
	// 协程计数器+1
	wg.Add(1)
	// 开启一个协程
	go test1()

	wg.Add(1)
	go test2()

	// 主线程的程序执行结束后就会退出，不会等待其他协程执行完毕
	for i := 0; i < 10; i++ {
		fmt.Println("main...", i)
		time.Sleep(time.Millisecond * 30)
	}

	// 等待协程执行完毕
	wg.Wait()
	fmt.Println("主线程执行完毕")
}
