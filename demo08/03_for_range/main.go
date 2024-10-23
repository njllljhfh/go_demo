package main

import "fmt"

func main() {
	var str = "你好glang"
	for k, v := range str {
		fmt.Printf("key=%v, value=%c\n", k, v)
	}
	fmt.Println("------------------------------------------")

	var arr = []string{"php", "java", "node", "golang"} // 切片
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }

	for k, v := range arr {
		fmt.Printf("key=%v, value=%s\n", k, v)
	}
	fmt.Println("------------------------------------------")

}
