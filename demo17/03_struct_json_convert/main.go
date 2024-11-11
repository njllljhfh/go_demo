package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id     int
	Gender string
	Name   string
}

type Class struct {
	Title    string
	Students []Student
}

func main() {
	c := Class{
		Title:    "001班",
		Students: make([]Student, 0),
	}

	for i := 1; i < 10; i++ {
		s := Student{
			Id:     i,
			Gender: "男",
			Name:   fmt.Sprintf("stu_%v", i),
		}
		c.Students = append(c.Students, s)
	}

	fmt.Println(c)
	fmt.Println("- - - - - - - - - - - - - - - - - -")

	// 嵌套结构体，转成json字符串
	if byteArr, err := json.Marshal(c); err != nil {
		fmt.Println(err)
	} else {
		strJson := string(byteArr)
		fmt.Println(strJson)
	}
	fmt.Println("- - - - - - - - - - - - - - - - - -")

	// json字符串 转 嵌套结构体
	str := `{"Title":"001班","Students":[{"Id":1,"Gender":"男","Name":"stu_1"},{"Id":2,"Gender":"男","Name":"stu_2"},{"Id":3,"Gender":"男","Name":"stu_3"},{"Id":4,"Gender":"男","Name":"stu_4"},{"Id":5,"Gender":"男","Name":"stu_5"},{"Id":6,"Gender":"男","Name":"stu_6"},{"Id":7,"Gender":"男","Name":"stu_7"},{"Id":8,"Gender":"男","Name":"stu_8"},{"Id":9,"Gender":"男","Name":"stu_9"}]}`
	c2 := &Class{}
	err := json.Unmarshal([]byte(str), c2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%#v\n", c2)
		fmt.Printf("%v\n", c.Title)
	}
}
