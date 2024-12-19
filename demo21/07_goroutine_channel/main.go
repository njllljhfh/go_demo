package main

import "fmt"

/*
   管道 channel:
       1.管道是引用类型
*/

func main() {

    // 定义管道
    ch := make(chan int, 3)

    // 给管道里面存数据
    ch <- 10
    ch <- 21
    ch <- 32

    // 从管道里面取数据
    a := <-ch
    fmt.Println(a) // 10

    <-ch

    c := <-ch
    fmt.Println(c) // 32
    fmt.Println("-----------------------")

    ch <- 56
    fmt.Printf("值:%v, 容量:%v, 长度:%v\n", ch, cap(ch), len(ch))
    fmt.Println("-----------------------")

    // 管道是引用类型
    ch1 := make(chan int, 4)
    ch1 <- 34
    ch1 <- 54
    ch1 <- 64

    // ch2 和 ch1 指向内存中的同一个管道实例
    ch2 := ch1
    ch2 <- 25

    <-ch1
    <-ch1
    <-ch1

    d := <-ch1
    fmt.Println(d) // 25
    fmt.Println("-----------------------")

    // 管道阻塞1
    // 在没有使用协程的情况下，管道数据满了，我们再向管道写数据会报错
    // ch6 := make(chan int, 1)
    // ch6 <- 34
    // ch6 <- 64 // 报错（死锁）： all goroutines are asleep - deadlock!
    // fmt.Println("-----------------------")

    // 管道阻塞2
    // 在没有使用协程的情况下，管道里没有数据了，我们再从管道取数据会报错
    // ch7 := make(chan string, 2)
    // ch7 <- "数据1"
    // ch7 <- "数据2"
    // m1 := <-ch7
    // m2 := <-ch7
    // m3 := <-ch7 // all goroutines are asleep - deadlock!
    // fmt.Println(m1, m2, m3)
    // fmt.Println("-----------------------")
}
