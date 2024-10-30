package main

import "fmt"

func main() {
	// 元素为slice类型的map

	var userinfo = make(map[string][]string)
	userinfo["hobby"] = []string{
		"吃饭",
		"睡觉",
		"敲代码",
	}
	fmt.Println(userinfo)
	fmt.Println("------------------------------------------")

	userinfo["work"] = []string{
		"php",
		"golang",
		"前端",
	}

	for k, v := range userinfo {
		for _, value := range v {
			fmt.Printf("k=%v, v=%v\n", k, value)
		}
	}
	fmt.Println("------------------------------------------")

}
