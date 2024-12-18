package main

import "fmt"

/*
	一个结构体实现多个接口
*/

type Animaler1 interface {
    SetName(string)
}

type Animaler2 interface {
    GetName() string
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
    // 结构体指针变量的方法集包括 结构体值接收者的方法集 和 结构体指针接收者的方法集
    d := &Dog{
        Name: "小黑",
    }

    // Dog实现 Animaler1 接口
    var d1 Animaler1 = d
    d1.SetName("小花狗狗")

    // Dog实现 Animaler2 接口
    var d2 Animaler2 = d
    fmt.Println(d2.GetName())
    fmt.Println("------------------------")
}
