package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	管道 channel:
	用管道和协程，实现素数统计 和 素数打印 并行执行
*/

// 向管道里存数字
// num: 自然数的个数
func putNum(num int, intChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= num; i++ {
		intChan <- i
	}
	close(intChan)
}

// primeNum 统计素数
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range intChan {
		tag := true
		for i := 2; i < v; i++ {
			if v%i == 0 {
				tag = false
				break
			}
		}
		if tag {
			primeChan <- v
		}
	}
	// 如果一个channel关闭了，就不能再向其发送数据了。
	// 1个协程执行完了，可能还有其他协程没有执行完，所以此处不能关闭 primeChan。
	// close(primeChan)

	exitChan <- true
}

// printPrime 打印找到的素数
func printPrime(primeChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	// 这里读取速度比写入速度快，读取时会等待下一个数据写入
	for v := range primeChan {
		fmt.Printf("%d 是素数\n", v)
	}
}

func main() {
	// 自然数的个数
	num := 120003
	// num := 100
	// 寻找素数的协程数
	workNum := 6

	// chanCap := int(float64(num) * 0.05)
	chanCap := 500
	fmt.Printf("管道容量 = %v\n", chanCap)

	startTime := time.Now().UnixMilli()

	/*
	   ch := make(chan int, 100)
	   容量参数 100 表示创建的是一个有缓冲通道，最多可以存储 100 个值。
	   特点：
	       - 发送操作不会立即阻塞，直到缓冲区满为止。
	       - 接收操作不会立即阻塞，直到缓冲区为空为止。
	       - 更适合需要临时存储数据、解耦生产者和消费者的场景。
	*/
	intChan := make(chan int, chanCap)
	primeChan := make(chan int, chanCap)
	exitChan := make(chan bool, workNum)

	/*
	   ch := make(chan int)
	   默认创建的是无缓冲通道。
	   特点：
	       - 发送操作 (ch <- value) 会阻塞，直到有另一个 Goroutine 从通道中接收值。
	       - 接收操作 (value := <-ch) 会阻塞，直到有另一个 Goroutine 向通道发送值。
	       - 更适合用于严格的 Goroutine 同步。*/
	// intChan := make(chan int)
	// primeChan := make(chan int)
	// exitChan := make(chan bool)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go putNum(num, intChan, wg)

	for i := 1; i <= workNum; i++ {
		wg.Add(1)
		go primeNum(intChan, primeChan, exitChan, wg)
	}

	wg.Add(1)
	go printPrime(primeChan, wg)

	// 所有协程都退出后，关闭未关闭的管道
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < workNum; i++ {
			<-exitChan
		}
		close(primeChan)
		close(exitChan) // 疑问：exitChan 要不要关闭，不关闭会怎么样
	}()

	wg.Wait()
	endTime := time.Now().UnixMilli()
	fmt.Printf("程序执行结束, 耗时 %v 毫秒", endTime-startTime)
}
