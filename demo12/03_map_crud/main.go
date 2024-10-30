package main

import "fmt"

func main() {
	// map类型的crud
	// 创建
	var userinfo = make(map[string]string)
	userinfo["username"] = "张三"
	userinfo["age"] = "20"
	fmt.Println(userinfo)
	fmt.Println("------------------------------------------")

	// 修改数据
	userinfo["username"] = "李四"
	fmt.Println(userinfo)
	fmt.Println("------------------------------------------")

	// 获取数据
	fmt.Println(userinfo["username"])
	fmt.Println("------------------------------------------")

	// 查找
	v, ok := userinfo["age"]
	fmt.Printf("v=%v, ok=%v\n", v, ok)
	v2, ok := userinfo["xxx"] // v2 是字符串的空值，ok=false
	fmt.Printf("v2=%v, ok=%v\n", v2, ok)
	fmt.Println("------------------------------------------")

	// 删除 map 里的key
	var userinfo2 = map[string]string{
		"username": "张三",
		"age":      "20",
		"sex":      "男",
		"height":   "180cm",
	}
	fmt.Println(userinfo2)
	delete(userinfo2, "sex")
	fmt.Println(userinfo2)
	fmt.Println("------------------------------------------")

}
