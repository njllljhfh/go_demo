package main

import "fmt"

// 自定义类型，添加方法
// 注意：非本包的类型不能添加方法，也就是说，不能给其他包的类型添加方法(例如，不能给int类型添加方法)
type MyInt int

func (m MyInt) PrintInfo() {
	fmt.Printf("m=%v, m的类型:%T\n", m, m)
}

func (m *MyInt) SetValue(v int) {
	*m = MyInt(v)
}

func main() {
	a := MyInt(20)
	a.PrintInfo()
	fmt.Println("- - - - - - - - - - - - - - - - - - - - ")

	b := MyInt(30)
	fmt.Printf("b=%v, b的类型:%T\n", b, b)
	b.SetValue(666)
	fmt.Printf("b=%v, b的类型:%T\n", b, b)
}
