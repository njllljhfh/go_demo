package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Sex    string
	height int
}

// 结构体方法和接收者
// 接收者为值类型
func (p Person) PrintInfo() {
	fmt.Printf("姓名：%v, 年龄：%v, 性别：%v, 身高：%v\n", p.Name, p.Age, p.Sex, p.height)
}

// 接收者为指针类型
func (p *Person) SetInfo(name string, age int) {
	p.Name = name
	p.Age = age
}

func main() {
	p1 := Person{
		Name: "张三",
		Age:  20,
		Sex:  "男",
	}
	p1.PrintInfo()
	fmt.Println("- - - - - - - - - - - - - - - - - - - - ")

	p1.SetInfo("李四", 34)
	p1.PrintInfo()
	fmt.Println("- - - - - - - - - - - - - - - - - - - - ")
}
