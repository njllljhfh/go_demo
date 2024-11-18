package main

import "fmt"

/*
接收者
*/

type Usber interface {
	start()
	stop()
}

// 手机
type Phone struct {
	Name string
}

// 指针接收者：结构体的值无法调用指针接收者的方法
func (p *Phone) start() {
	fmt.Printf("%v, 启动\n", p.Name)
}

func (p *Phone) stop() {
	fmt.Printf("%v, 关机\n", p.Name)
}

func main() {

	// 值
	// p1 := Phone{
	// 	Name: "小米手机",
	// }
	// var p2 Usber = p1 // 报错：cannot use p1 (variable of type Phone) as Usber value in variable declaration: Phone does not implement Usber (method start has pointer receiver)
	// p2.start()

	// 指针
	var p3 = &Phone{
		Name: "苹果手机",
	}
	var p4 Usber = p3
	p4.start()
	fmt.Println("------------------------")

}
