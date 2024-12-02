package main

import (
	"fmt"
	"io"
	"os"
)

/*
   复制文件操作
*/

func mCopyFile(srcFileName string, dstFileName string) error {
	// TODO: 这里在函数内定义 异常，是有问题的，go函数的return分两步！！！
	// 	   第一步：将 finalErr 的值 赋值 给 函数底层维护的返回值
	//     第二步：将 函数底层维护的返回值 返回给调用方
	// 	   重点：defer函数执行的时机是在上述的第一步之后，在上述的第二步之前！！！
	var finalErr error = nil

	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("未知异常:%v\n", e)
		}
	}()

	sFile, finalErr := os.Open(srcFileName)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			finalErr = fmt.Errorf("关闭src文件失败: %v\n", err)
		} else {
			fmt.Println("关闭src文件成功")
		}
	}(sFile)

	dFile, finalErr := os.OpenFile(dstFileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			finalErr = fmt.Errorf("关闭dst文件失败: %v\n", err)
		} else {
			fmt.Println("关闭dst文件成功")
		}
	}(dFile)

	tempSlice := make([]byte, 32)
	for {
		// 读取
		n, err1 := sFile.Read(tempSlice)
		if err1 == io.EOF {
			break
		}
		if err1 != nil {
			return fmt.Errorf("读取文件流失败: %v\n", err1)
		}

		// 写入
		_, finalErr = dFile.Write(tempSlice[:n])
	}

	return finalErr
}

func main() {
	src := "./test.txt"
	dst := "./copyFile.txt"
	err := mCopyFile(src, dst)
	if err != nil {
		fmt.Println("报错啦！！！ ", err)
		return
	}
	fmt.Println("复制文件成功")
}
