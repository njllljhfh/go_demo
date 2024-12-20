package main

import (
    "fmt"
    "sync"
    "time"
)

/*
   单向管道
*/

// fn1 写数据
// 参数
// ch:只写管道
func fn1(ch chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 10; i++ {
        ch <- i
        fmt.Printf("[写入]数据 %v 成功\n", i)
        time.Sleep(time.Millisecond * 50)
    }
    close(ch)
}

// fn2 读数据
// 参数
// ch:只读管道
func fn2(ch <-chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    // 使用协程的情况下，写的快，读取的慢，管道满了也不会报错
    for v := range ch {
        fmt.Printf("[读取]数据 %v 成功\n", v)
        time.Sleep(time.Millisecond * 50)
    }
}

func main() {
    wg := &sync.WaitGroup{}

    // 双向管道
    ch := make(chan int, 2)

    wg.Add(1)
    go fn1(ch, wg)

    wg.Add(1)
    go fn2(ch, wg)

    wg.Wait()

    fmt.Println("-----------------------")
}
