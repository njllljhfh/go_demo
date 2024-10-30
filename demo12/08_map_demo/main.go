package main

import (
	"fmt"
	"strings"
)

func main() {
	// 统计字符串中，每个单词出现的次数
	str := "how do you do"
	strSlice := strings.Split(str, " ")
	fmt.Println(strSlice)

	strMap := make(map[string]int)
	for _, v := range strSlice {
		strMap[v]++ // 小技巧；没有key，就会创建key，有key，就会修改value
	}
	fmt.Println(strMap)
}
