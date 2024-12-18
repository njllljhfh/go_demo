package main

import "fmt"

/*
接收者
*/

type Usber interface {
    start()
    stop()
}

// Phone 手机
type Phone struct {
    Name string
}

// 值接收者：如果结构体中的方法是值接收者，那么这个方法可以通过实例调用，也可以通过指针调用
func (p Phone) start() {
    fmt.Printf("%v, 启动\n", p.Name)
}

func (p Phone) stop() {
    fmt.Printf("%v, 关机\n", p.Name)
}

func main() {

    // 值
    p1 := Phone{
        Name: "小米手机",
    }

    var p2 Usber = p1
    p2.start()

    // 指针
    var p3 = &Phone{
        Name: "苹果手机",
    }

    var p4 Usber = p3
    p4.start()
    fmt.Println("------------------------")

}
