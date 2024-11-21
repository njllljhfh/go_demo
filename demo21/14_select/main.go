package main

import (
	"fmt"
)

/*
	select 多路复用

    在某些场景下，我们需要溶蚀动多个通道接受数据，这时可以使用 select 多路复用。

	select 的执行逻辑
		- 非阻塞随机选择：
			- select 会随机选择一个可以执行的 case 分支。如果有多个分支满足条件，Go 会随机挑选其中一个。
			- 如果所有 case 都不能执行且存在 default 分支，则执行 default 分支。

		- 阻塞等待：
			- 如果所有 case 分支都不满足条件，且没有 default 分支，则 select 会阻塞，直到某个 case 能够执行。
*/

func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	// close(intChan) // 不要关闭channel，否则会导致 下面一直循环下去

	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}
	// close(stringChan) // 不要关闭channel，否则会导致 下面一直循环下去

	// 注意：使用select来获取 channel 里面的数据时 不需要关闭 channel
	for {
		select {
		case v := <-intChan:
			fmt.Printf("从 intChan 读取数据%d\n", v)
		case s := <-stringChan:
			fmt.Printf("从 stringChan 读取数据%v\n", s)
		default:
			fmt.Println("数据获取完毕")
			return
		}
	}
}
