package main

import (
    "fmt"
    "sync"
    "time"
)

/*
	读写锁
		读锁：共享锁
		写锁：排它锁
*/

var wg sync.WaitGroup

// 读写锁
var mutex sync.RWMutex

// 写的方法
func write() {
    // 写锁：当一个goroutine进行写操作时，其他goroutine既不能读，也不能写。
    mutex.Lock() // 写锁
    fmt.Println("对资源 a 执行 写操作")
    time.Sleep(time.Second * 1) // 模拟耗时操作
    mutex.Unlock()
    wg.Done()
}

// 读操作
func read() {
    // 读锁：当一个goroutine进行读操作时，其他goroutine可以读，不可以写。
    mutex.RLock() // 读锁
    fmt.Println("---对资源 a 执行 读操作")
    time.Sleep(time.Second * 1)
    mutex.RUnlock()
    wg.Done()
}

func main() {

    // 开启协程执行读操作
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go read()
    }

    // 开启协程执行写操作
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go write()
    }

    wg.Wait()
}
