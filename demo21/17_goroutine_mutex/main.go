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

	互斥锁:解决资源竞争
*/

var count = 0
var wg sync.WaitGroup

// 互斥锁
var mutex sync.Mutex

func test() {
    // 加锁
    mutex.Lock()
    count++
    fmt.Printf("the count is %d\n", count)
    time.Sleep(time.Millisecond)
    // 释放锁
    mutex.Unlock()
    wg.Done()
}

func main() {
    for r := 0; r < 20; r++ {
        wg.Add(1)
        go test()
    }

    wg.Wait()
    fmt.Printf("main() --- the count is %d\n", count)
}
