package main

import (
	"fmt"
	"os"
)

/*
   复制文件操作
*/

func Copy(srcFileName string, dstFileName string) error {
	// 读取
	byteSlice, err := os.ReadFile(srcFileName)
	if err != nil {
		return fmt.Errorf("读取文件失败: %v\n", err)
	}

	// 写入
	err = os.WriteFile(dstFileName, byteSlice, 0666)
	if err != nil {
		return fmt.Errorf("复制文件失败: %v\n", err)
	}

	return nil
}

func main() {
	src := "./demo23/test.txt"
	dst := "./demo23/07_file_copy/copyFile.txt"
	err := Copy(src, dst)
	if err != nil {
		fmt.Println("复制文件失败！！！")
		return
	}
	fmt.Println("复制文件成功")
}
