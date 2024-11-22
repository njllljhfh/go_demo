package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	资源竞争
	go build -race main.go 编译后运行查看
	-race 查看编译后的可执行程序在运行的时候有没有竞争
*/

var count = 0
var wg sync.WaitGroup

func test() {
	count++
	fmt.Printf("the count is %d\n", count)
	time.Sleep(time.Millisecond)
	wg.Done()
}

func main() {
	// 资源竞争
	for r := 0; r < 20; r++ {
		wg.Add(1)
		go test()
	}

	wg.Wait()
	// 此处count 可能不等于 20
	fmt.Printf("main() --- the count is %d\n", count)
}
