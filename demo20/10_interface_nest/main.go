package main

import "fmt"

/*
	接口嵌套
*/

type AInterface interface {
	SetName(string)
}

type BInterface interface {
	GetName() string
}

// 接口嵌套
type Animaler interface {
	AInterface
	BInterface
}

type Dog struct {
	Name string
}

func (d *Dog) SetName(name string) {
	d.Name = name
}

func (d Dog) GetName() string {
	return d.Name
}

func main() {

	d := &Dog{
		Name: "小黑",
	}

	var d1 Animaler = d
	d1.SetName("小花狗狗")
	fmt.Println(d1.GetName())
	fmt.Println("------------------------")

}
