package main

import (
	"fmt"
	"time"
)

/*
	协程
*/

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test...", i)
		time.Sleep(time.Millisecond * 10)
	}
}

func main() {
	// 开启一个协程
	go test()

	// 主线程的程序执行结束后就会退出，不会等待其他协程执行完毕
	for i := 0; i < 10; i++ {
		fmt.Println("main...", i)
		time.Sleep(time.Millisecond * 50)
	}
}
