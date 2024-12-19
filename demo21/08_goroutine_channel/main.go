package main

import "fmt"

/*
   管道 channel:
	遍历管道
*/

func main() {

    ch1 := make(chan int, 10)
    for i := 1; i <= 10; i++ {
        ch1 <- i
    }

    // for range 遍历管道，管道写入完成要关闭，否则报错：all goroutines are asleep - deadlock!
    // 关闭管道
    close(ch1)

    // 管道没有 key
    for v := range ch1 {
        fmt.Println(v)
    }

    fmt.Println("-----------------------")

    // 通过 for 循环遍历管道，管道可以不关闭
    ch2 := make(chan int, 10)
    for i := 1; i <= 10; i++ {
        ch2 <- i
    }
    for j := 1; j <= 10; j++ {
        fmt.Println(<-ch2)
    }

}
