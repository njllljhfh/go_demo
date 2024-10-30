package main

import "fmt"

/*
	值类型：基本数据类型, 数组
	引用类型：slice, map
*/

func main() {
	var userinfo1 = make(map[string]string)
	userinfo1["username"] = "张三"
	userinfo1["age"] = "20"

	userinfo2 := userinfo1
	fmt.Println(userinfo1)
	fmt.Println(userinfo2)

	userinfo2["username"] = "李四"
	fmt.Println(userinfo1) // 修改了userinfo2，userinfo1也会跟着修改
	fmt.Println(userinfo2)
	fmt.Println("------------------------------------------")

}
