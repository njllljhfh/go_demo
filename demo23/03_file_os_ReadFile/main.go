package main

import (
	"fmt"
	"os"
)

/*
   读取文件操作: ioutil(go 1.16版本后,已经弃用了)，现在用 os包的相关方法。
*/

func main() {
	// ioutil.ReadFile()  go 1.16版本后,已经弃用了
	// 一次读取文件的全部内容
	byteSlice, err := os.ReadFile("./demo23/test.txt")
	if err != nil {
		fmt.Println("读取文件失败:", err)
		return
	}
	fmt.Println(string(byteSlice))
}
