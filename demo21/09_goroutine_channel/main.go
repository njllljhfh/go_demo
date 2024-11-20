package main

import (
    "fmt"
    "sync"
    "time"
)

/*
   管道 channel:
*/

// fn1 写数据
func fn1(ch chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 10; i++ {
        ch <- i
        fmt.Printf("[写入]数据 %v 成功\n", i)
        time.Sleep(time.Millisecond * 1)
    }
    close(ch)
}

// fn2 读数据
func fn2(ch chan int, wg *sync.WaitGroup) {
    defer wg.Done()
    // 使用协程的情况下，写的快，读取的慢，管道满了也不会报错
    for v := range ch {
        fmt.Printf("[读取]数据 %v 成功\n", v)
        time.Sleep(time.Millisecond * 1000)
    }
}

func main() {
    fmt.Println("使用协程的情况下，写的快，读取的慢，管道满了也不会报错")

    // var wg *sync.WaitGroup
    // wg = new(sync.WaitGroup)
    wg := &sync.WaitGroup{}

    ch := make(chan int, 2)

    wg.Add(1)
    go fn1(ch, wg)
    wg.Add(1)
    go fn2(ch, wg)
    wg.Wait()

    fmt.Println("-----------------------")
}
