package main

import (
	"fmt"
	"reflect"
)

/*
   反射：修改结构体属性
*/

/*
不要乱用反射
    反射式一个强大并富有表现力的工具，能让我们写出更灵活的代码。但是反射不应该被滥用，主要有以下原因。
        1. 基于反射的代码是极其脆弱的，反射中的类型错误会在运行时才会引发panic，那很可能是在代码写完的很长时间之后。
        2. 大量使用反射的代码通常难以理解。
*/

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func (s *Student) GetInfo() string {
	str := fmt.Sprintf("姓名:%v, 年龄:%v, 成绩:%v\n", s.Name, s.Age, s.Score)
	return str
}

// 反射修改结构体属性
func reflectChangeStruct(s interface{}) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	if !(t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct) {
		fmt.Println("传入的参数不是一个结构体指针类型")
		fmt.Println("------------------------")
		return
	}

	// 修改结构体属性的值
	name := v.Elem().FieldByName("Name")
	name.SetString("小李")

	age := v.Elem().FieldByName("Age")
	age.SetInt(22)
}

func main() {
	stu1 := Student{
		Name:  "小明",
		Age:   15,
		Score: 98,
	}

	// reflectChangeStruct(stu1)

	// a := 1
	// reflectChangeStruct(&a)

	// 通过反射修改结构体属性的值时，要传入 结构体的指针（因为结构体是值类型）
	reflectChangeStruct(&stu1)
	fmt.Printf("%#v\n", stu1) // main.Student{Name:"小李", Age:22, Score:98}
}
