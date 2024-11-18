package main

import "fmt"

/*
	类型断言
*/

// 写法1
func MyPrint1(x interface{}) {
	if v, ok := x.(string); ok {
		fmt.Printf("类型为 string, 值为：%v\n", v)
	} else if v, ok := x.(int); ok {
		fmt.Printf("类型为 int, 值为：%v\n", v)
	} else if v, ok := x.(bool); ok {
		fmt.Printf("类型为 bool, 值为：%v\n", v)
	} else {
		fmt.Println("类型断言失败，未知类型")
	}
}

func MyPrint2(x interface{}) {
	// x.(type) 语句只能用在 switch 语句中，不能用在其他地方。
	switch v := x.(type) {
	case string:
		fmt.Printf("类型为 string, 值为：%v\n", v)
	case int:
		fmt.Printf("类型为 int, 值为：%v\n", v)
	case bool:
		fmt.Printf("类型为 bool, 值为：%v\n", v)
	default:
		fmt.Println("类型断言失败，未知类型")
	}
}

// 写法2
func main() {
	var a interface{}
	a = "你好"
	// a = 123
	v, ok := a.(string)
	if ok {
		fmt.Printf("类型断言成功，值为：%v\n", v)
	} else {
		fmt.Println("类型断言失败")
	}
	fmt.Println("-------------------------------")

	MyPrint1("你好")
	fmt.Println("-------------------------------")

	MyPrint2(101)
}
