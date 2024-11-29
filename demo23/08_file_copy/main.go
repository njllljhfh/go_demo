package main

import (
    "fmt"
    "io"
    "os"
)

/*
   复制文件操作
*/

func CopyFile(srcFileName string, dstFileName string, fenMu int) (deferErr error) {
    // 处理未知异常
    defer func() {
        if err := recover(); err != nil {
            fmt.Printf("未知异常:%v\n", err)
        }
    }()

    sFile, deferErr := os.Open(srcFileName)
    if deferErr != nil {
        return
    }
    // defer sFile.Close()
    defer func(f *os.File) {
        err := f.Close()
        if err != nil {
            deferErr = fmt.Errorf("关闭src文件失败: %v\n", err)
        } else {
            fmt.Println("关闭src文件成功")
        }
    }(sFile)

    dFile, deferErr := os.OpenFile(dstFileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
    if deferErr != nil {
        return
    }
    // defer dFile.Close()
    defer func(f *os.File) {
        err := f.Close()
        // err = fmt.Errorf("模拟错误---关闭dst文件失败") // TODO: 模拟关闭报错
        if err != nil {
            deferErr = fmt.Errorf("关闭dst文件失败: %v\n", err)
        } else {
            fmt.Println("关闭dst文件成功")
        }
    }(dFile)

    a := 1 / fenMu
    fmt.Printf("a = %v\n", a)

    tempSlice := make([]byte, 32)
    for {
        // 读取
        n, err := sFile.Read(tempSlice)
        if err == io.EOF {
            break
        }
        if err != nil {
            deferErr = fmt.Errorf("读取文件流失败: %v\n", err)
            return
        }

        // 写入
        _, err = dFile.Write(tempSlice[:n])
        if err != nil {
            deferErr = fmt.Errorf("写入文件流失败: %v\n", err)
            return
        }
    }

    return
}

func main() {
    src := "../test.txt"
    dst := "./copyFile.txt"
    fenMu := 1
    err := CopyFile(src, dst, fenMu)
    if err != nil {
        fmt.Println("报错啦！！！ ", err)
        return
    }
    fmt.Println("复制文件成功")
}
