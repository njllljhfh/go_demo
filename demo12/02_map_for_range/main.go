package main

import "fmt"

func main() {
	// 遍历
	var userinfo = map[string]string{
		"username": "张三",
		"age":      "20",
		"sex":      "男",
	}
	for k, v := range userinfo {
		fmt.Printf("k=%v, v=%v\n", k, v)
	}
}
