package main

import "fmt"

// map 是无序的，key不能重复，value可以重复
// map 是引用类型，默认值是nil，不能直接赋值，需要使用make创建

func main() {
	// 1.make创建map
	var userinfo = make(map[string]string)
	userinfo["unsername"] = "张三"
	userinfo["age"] = "20"
	userinfo["sex"] = "男"
	fmt.Println(userinfo)
	fmt.Println(userinfo["unsername"])
	fmt.Println("------------------------------------------")

	// 2.字面量创建map
	var userinfo2 = map[string]string{
		"username": "张三",
		"age":      "20",
		"sex":      "男",
	}
	fmt.Println(userinfo2)
	fmt.Println("------------------------------------------")

	// 3
	userinfo3 := map[string]string{
		"username": "张三",
		"age":      "20",
		"sex":      "男",
	}
	fmt.Println(userinfo3)
	fmt.Println("------------------------------------------")

}
