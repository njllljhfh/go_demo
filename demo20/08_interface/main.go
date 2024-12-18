package main

import "fmt"

/*
	接口中方法的参数和返回值
*/

type Animaler interface {
    SetName(string)
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

type Cat struct {
    Name string
}

func (c *Cat) SetName(name string) {
    c.Name = name
}

func (c Cat) GetName() string {
    return c.Name
}

func main() {

    d := &Dog{
        Name: "小黑",
    }
    var d1 Animaler = d
    fmt.Println(d1.GetName())
    d1.SetName("小白")
    fmt.Println(d1.GetName())
    fmt.Println("------------------------")

    c := &Cat{
        Name: "小花",
    }
    var c1 Animaler = c
    fmt.Println(c1.GetName())
    c1.SetName("小红")
    fmt.Println(c1.GetName())
}
