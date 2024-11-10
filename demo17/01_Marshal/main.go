package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID     int
	Gender string
	name   string // 私有属性不能被json包访问
	Sno    string
}

func main() {
	s1 := Student{
		ID:     12,
		Gender: "男",
		name:   "张三",
		Sno:    "s0001",
	}

	fmt.Printf("%#v\n", s1)

	// 结构体 转 json 字符串
	jsonByte, _ := json.Marshal(s1) // jsonByte 的类型是 []byte
	// fmt.Printf("%v\n", jsonByte)
	jsonStr := string(jsonByte)
	fmt.Printf("%v\n", jsonStr)
	fmt.Println("------------------------------------------")

	// json 字符串 转 结构体
	str := `{"ID":12,"Gender":"男","name":"张三","Sno":"s0001"}`
	s2 := Student{}
	err := json.Unmarshal([]byte(str), &s2)
	if err != nil {
		fmt.Println("json.Unmarshal err:", err)
	}
	fmt.Printf("%#v\n", s2) // main.Student{ID:12, Gender:"男", name:"", Sno:"s0001"
	fmt.Println(s2.name)    // name 是空字符串
}
