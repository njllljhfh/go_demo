package main

import (
	"encoding/json"
	"fmt"
)

//结构体标签(实际项目中用的很多)

type Student struct {
	ID     int    `json:"id"`
	Gender string `json:"gender"`
	Name   string `json:"name"`
	Sno    string `json:"xxx"`
}

func main() {
	s := Student{
		ID:     12,
		Gender: "男",
		Name:   "张三",
		Sno:    "s0001",
	}

	byteArr, _ := json.Marshal(s)
	jsonStr := string(byteArr)
	fmt.Println(jsonStr)
	fmt.Println("------------------------------------------")

	// json 字符串 转 结构体(json中的小写字母会根据 结构体中的标签 转成大写字母)
	str := `{"id":12,"gender":"男","name":"张三","xxx":"s0001"}`
	s2 := &Student{}
	json.Unmarshal([]byte(str), &s2)
	fmt.Printf("%#v\n", s2)
}
