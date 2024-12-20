package main

import "fmt"

/*
单向管道
*/

func main() {

    // 默认情况下，管道是双向的
    ch1 := make(chan int, 2)
    ch1 <- 19
    ch1 <- 12
    m1 := <-ch1
    m2 := <-ch1
    fmt.Println(m1, m2)
    fmt.Println("-----------------------")

    // 只写管道
    ch2 := make(chan<- int, 2)
    ch2 <- 19
    ch2 <- 12
    // <-ch2 // invalid operation: cannot receive from send-only channel ch2 (variable of type chan<- int)
    fmt.Println("-----------------------")

    // 只读管道
    // ch3 := make(<-chan int, 2)
    // ch3 <- 23 // invalid operation: cannot send to receive-only channel ch3 (variable of type <-chan int)
}
