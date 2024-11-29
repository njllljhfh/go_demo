package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
   写入文件操作:
        file.WriteString()
        file.Write()
*/

func main() {
	// 0666：是文件的权限（是用八进制表示的）
	file, err := os.OpenFile("./demo23/04_file_io_write/textWrite.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	defer func(file *os.File) {
		err := file.Close() // 关闭文件
		if err != nil {
			fmt.Println("文件关闭失败:", err)
		}
	}(file)

	if err != nil {
		fmt.Println("文件打开失败:", err)
	}

	for i := 0; i < 10; i++ {
		// 写入文件
		_, err := file.WriteString("直接写入的字符串数据" + strconv.Itoa(i) + "\r\n")
		if err != nil {
			fmt.Printf("写入文件失败: %v\n", err)
			return
		}
	}

	str := "转换为字节切片后再写入数据\r\n"
	_, _ = file.Write([]byte(str))
}
