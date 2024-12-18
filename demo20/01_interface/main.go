package main

import "fmt"

// Usber 如果接口里面有方法的话，必须要通过结构体或通过自定义类型来实现这个接口
type Usber interface {
    start()
    stop()
}

// Phone 手机
// 要实现接口，必须要实现接口里面的所有方法
type Phone struct {
    Name string
}

func (p Phone) start() {
    fmt.Printf("%v, 启动\n", p.Name)
}
func (p Phone) stop() {
    fmt.Printf("%v, 关机\n", p.Name)
}

// Camera 相机
type Camera struct {
}

func (c Camera) start() {
    fmt.Printf("相机启动\n")
}
func (c Camera) stop() {
    fmt.Printf("相机关机\n")
}

func (c Camera) run() {
    fmt.Println("相机运行")
}

func main() {
    p := Phone{
        Name: "华为手机",
    }

    var p1 Usber
    p1 = p
    p1.start()
    p1.stop()
    fmt.Println("--------------------------------")

    var c1 Usber
    c := Camera{}
    c1 = c
    c1.start()
    c1.stop()

    // c1.run()  // 无法调用接口中没有的方法
    c.run() // 用实例调用
}
