package main

import "fmt"

/*
	一个结构体中可以嵌套另一个结构体或结构体指针。
*/

// 父结构体
type Animal struct {
	Name string
}

func (a Animal) run() {
	fmt.Printf("%v 在运动\n", a.Name)
}

func (a *Animal) eat() {
	fmt.Printf("%v 在吃饭\n", a.Name)
}

// 子结构体
type Dog struct {
	Age     int
	*Animal // 嵌套结构体指针
}

func (d Dog) wang() {
	fmt.Printf("%v 在旺旺\n", d.Name)
}

func main() {
	d := Dog{
		Age: 20,
		Animal: &Animal{
			Name: "阿奇",
		},
	}
	d.run()
	d.wang()
	d.eat()
}
