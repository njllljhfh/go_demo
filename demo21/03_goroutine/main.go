package main

import (
    "fmt"
    "runtime"
)

func main() {
    // 获取当前计算机上面的CPU数量
    cpuNum := runtime.NumCPU()
    fmt.Println("cpuNum=", cpuNum)

    // 可以自己设置用几个cpu
    runtime.GOMAXPROCS(cpuNum - 1)
    fmt.Println("ok")
}
