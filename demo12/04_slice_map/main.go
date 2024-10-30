package main

import "fmt"

func main() {
	// 元素为map类型的切片
	var userinfo = make([]map[string]string, 3, 3)
	fmt.Println(userinfo[0])        // map[],
	fmt.Println(userinfo[0] == nil) // true, map未初始化，默认值为nil

	if userinfo[0] == nil {
		userinfo[0] = make(map[string]string)
		userinfo[0]["username"] = "张三"
		userinfo[0]["age"] = "20"
		userinfo[0]["height"] = "180cm"
	}

	if userinfo[1] == nil {
		userinfo[1] = make(map[string]string)
		userinfo[1]["username"] = "李四"
		userinfo[1]["age"] = "22"
		userinfo[1]["height"] = "170cm"
	}
	fmt.Println(userinfo)
	fmt.Println("------------------------------------------")

	// 遍历
	for _, v := range userinfo {
		for k, v := range v {
			fmt.Printf("k=%v, v=%v\n", k, v)
		}
		fmt.Println("---")
	}

}
