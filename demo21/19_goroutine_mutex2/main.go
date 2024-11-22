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

var wg sync.WaitGroup

var m = make(map[int]int)

var mutex sync.Mutex

func test(num int) {
	// 阶乘
	sum := 1
	for i := 1; i <= num; i++ {
		sum *= i
	}

	mutex.Lock()
	m[num] = sum
	mutex.Unlock()

	fmt.Printf("key=%v, value=%v\n", num, sum)
	time.Sleep(time.Millisecond)
	wg.Done()
}

func main() {
	// 资源竞争
	for r := 0; r < 20; r++ {
		wg.Add(1)
		go test(r)
	}

	wg.Wait()
}
