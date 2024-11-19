package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 获取当前计算机上面的CPU数量
	cupNum := runtime.NumCPU()
	fmt.Println("cupNum=", cupNum)

	// 可以自己设置只用多个cpu
	runtime.GOMAXPROCS(cupNum - 1)
	fmt.Println("ok")
}
