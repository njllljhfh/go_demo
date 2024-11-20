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
		time.Sleep(time.Millisecond * 500)
	}
	close(ch)
}

/*
以下解释来自ChatGPT

    为什么会报死锁错误？
    for range 在遍历管道时，会一直等待数据的到来。如果没有新的数据写入，并且管道没有关闭，for range 会阻塞，导致程序无法继续运行。

    管道未关闭的情况
    当 for range 遍历到管道的最后一个元素后：

    它会尝试从管道中继续读取。
    如果没有新数据且管道未关闭，for range 会进入阻塞状态，等待新的数据写入。
    如果此时所有的 goroutine 都已经完成运行，而主 goroutine 阻塞在 for range 中，Go 会检测到所有的 goroutine 都处于“休眠”状态并报 deadlock 错误。
*/

// fn2 读数据
func fn2(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	// 这里读取速度比写入速度快，读取时会等待下一个数据写入
	for v := range ch {
		fmt.Printf("[读取]数据 %v 成功\n", v)
		time.Sleep(time.Millisecond * 10)
	}
}

func main() {
	// var wg *sync.WaitGroup
	// wg = new(sync.WaitGroup)
	wg := &sync.WaitGroup{}

	ch := make(chan int, 10)

	wg.Add(1)
	go fn1(ch, wg)
	wg.Add(1)
	go fn2(ch, wg)
	wg.Wait()

	fmt.Println("-----------------------")
}
