package main

import (
	"fmt"
	"os"
)

/*
   读取文件操作: ioutil(go 1.16版本后,已经弃用了)，现在用 os包的相关方法。
*/

func main() {
	// 写入文件数据
	str := "你好nice！！！\r\n"
	fileName := "./demo23/06_file_os_WriteFile/testWrite.txt"
	err := os.WriteFile(fileName, []byte(str), 0666)
	if err != nil {
		fmt.Println("写入文件失败:", err)
	}
}
