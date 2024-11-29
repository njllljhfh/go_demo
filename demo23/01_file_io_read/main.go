package main

import (
	"fmt"
	"io"
	"os"
)

/*
   读取文件操作: file.Read()
*/

func main() {
	// [只读]的方式打开文件
	file, err := os.Open("./demo23/test.txt")

	// 必须要关闭文件流
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("关闭文件报错:", err)
		}
	}(file)

	if err != nil {
		fmt.Println("打开文件报错:", err)
		return
	}

	// 操作文件
	fmt.Println(file) // file 是一个指针(*os.File)

	// 读取文件内容
	// allSlice := make([]byte, 0)
	var allSlice []byte
	tempSlice := make([]byte, 128)
	for {
		n, err := file.Read(tempSlice) // 最多只读取 128 个字节

		// 文件读完了 会抛出 io.EOF
		if err == io.EOF {
			fmt.Println("文件读取完毕!")
			break
		}
		if err != nil {
			fmt.Println("读取文件内容失败")
			return
		}

		fmt.Printf("读取到了 %v 个字节\n", n)
		allSlice = append(allSlice, tempSlice[:n]...) // tempSlice每次只会从左到右覆盖n个字节，索引 >= n 的字节不会被清空
	}
	fmt.Println(string(allSlice))
}
