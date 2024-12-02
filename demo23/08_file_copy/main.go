package main

import (
    "fmt"
    "io"
    "os"
)

/*
   复制文件操作
*/

func CopyFile(srcFileName string, dstFileName string, fenMu int) (err error) {
    // 处理未知异常
    defer func() {
        if err1 := recover(); err1 != nil {
            fmt.Printf("未知异常:%v\n", err1)
        }
    }()

    sFile, err2 := os.Open(srcFileName)
    if err2 != nil {
        return err2
    }
    // defer sFile.Close()
    defer func(f *os.File) {
        err3 := f.Close()
        if err3 != nil {
            err = fmt.Errorf("关闭src文件失败: %v\n", err3)
        } else {
            fmt.Println("关闭src文件成功")
        }
    }(sFile)

    dFile, err4 := os.OpenFile(dstFileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
    if err4 != nil {
        return err4
    }
    // defer dFile.Close()
    defer func(f *os.File) {
        err5 := f.Close()
        // err5 = fmt.Errorf("模拟错误---关闭dst文件失败") // TODO: 模拟关闭报错
        if err5 != nil {
            err = fmt.Errorf("关闭dst文件失败: %v\n", err5)
        } else {
            fmt.Println("关闭dst文件成功")
        }
    }(dFile)

    a := 1 / fenMu // TODO: 模拟未知错误
    fmt.Printf("a = %v\n", a)

    tempSlice := make([]byte, 32)
    for {
        // 读取
        n, err6 := sFile.Read(tempSlice)
        if err6 == io.EOF {
            break
        }
        if err6 != nil {
            return fmt.Errorf("读取文件流失败: %v\n", err6)
        }

        // 写入
        _, err6 = dFile.Write(tempSlice[:n])
        if err6 != nil {
            return fmt.Errorf("写入文件流失败: %v\n", err6)
        }
    }

    return
}

func main() {
    src := "../test.txt"
    dst := "./copyFile.txt"
    // fenMu := 0 // TODO: 模拟未知错误
    fenMu := 1
    err := CopyFile(src, dst, fenMu)
    if err != nil {
        fmt.Println("报错啦！！！ ", err)
        return
    }
    fmt.Println("复制文件成功")
}
