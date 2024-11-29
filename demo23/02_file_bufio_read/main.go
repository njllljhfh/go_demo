package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
   读取文件操作: bufio
*/

func main() {
	// [只读]的方式打开文件
	file, err := os.Open("./demo23/test.txt")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("关闭文件失败:", err)
		}
	}(file)

	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}

	// bufio 读取文件
	reader := bufio.NewReader(file)
	fileStr := ""
	for {
		// 用这种方式读取时，当err为io.EOF时，str也可能是有数据的（即文件的最后一行数据）
		str, err := reader.ReadString('\n') // 表示一次读取一行
		if err == io.EOF {                  // 读取完毕
			fileStr += str
			fmt.Printf("读取文件结束, str=%v\n", str)
			fmt.Println("读取文件结束:", err)
			break
		}
		if err != nil {
			fmt.Println("读取异常:", err)
			return
		}
		fileStr += str
	}
	fmt.Println(fileStr)
}
