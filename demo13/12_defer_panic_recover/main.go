package main

import (
    "errors"
    "fmt"
)

func readFile(fileName string) error {
    if fileName == "main.go" {
        return nil
    } else {
        return errors.New("读取文件失败")
    }
}

func myFn() {
    // 相当于其他语言的 try catch finally
    defer func() {
        // 捕获 panic, 从 panic 中恢复
        if unknownErr := recover(); unknownErr != nil {
            fmt.Printf("给管理员发送邮件: %v\n", unknownErr)
        }
    }()

    err := readFile("xxx.go")
    if err != nil {
        panic(err)
    }
}

func main() {
    myFn()
    fmt.Println("继续执行...")
}
