package main

import (
    "fmt"
    "time"
)

/*
 */

func sayHello() {
    for i := 0; i < 10; i++ {
        time.Sleep(time.Millisecond * 50)
        fmt.Printf("hello world %d\n", i)
    }
}

func test() {
    // 使用 defer + recover 处理异常。即使 test 执行出错，不会影响 sayHello 的执行。
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("test()函数放生错误", err)
        }
    }()

    // 声明一个map
    var myMap map[int]string
    myMap[0] = "golang" // error: assignment to entry in nil map
}

func main() {
    go sayHello()
    go test()

    // 防止主进程退出
    time.Sleep(time.Second)
}
