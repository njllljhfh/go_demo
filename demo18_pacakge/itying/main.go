package main

/*
	包名为main的包为应用程序的入口包，编译后会生成可执行文件，而编译不包含 main 包的源代码文件则不会生成可执行文件
	一个项目里只能有一个main包，main包中只能有一个main函数
*/

import (
	"fmt"
	"itying/calc"
)

// go mod init 项目根目录名称 (生成go.mod文件)

func main() {
	sum := calc.Add(10, 2)
	fmt.Println(sum)

	fmt.Printf("Age = %d\n", calc.Age)
}
