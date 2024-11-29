package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
   写入文件操作: bufio
*/

func main() {
	// 0666：是文件的权限（是用八进制表示的）
	file, err := os.OpenFile("./demo23/05_file_bufio_write/textWrite.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	defer func(file *os.File) {
		err := file.Close() // 关闭文件
		if err != nil {
			fmt.Println("文件关闭失败:", err)
		}
	}(file)

	if err != nil {
		fmt.Println("文件打开失败:", err)
	}

	writer := bufio.NewWriter(file)

	for i := 0; i < 10; i++ {
		// 将数据先写入缓存中
		_, _ = writer.WriteString("你好-哦买嘎嘎..." + strconv.Itoa(i) + "\r\n")
		// 将缓存中的数据写入文件
		_ = writer.Flush()
	}

	str := "转换为字节切片后再写入数据\r\n"
	_, _ = writer.Write([]byte(str))
	_ = writer.Flush()
}
