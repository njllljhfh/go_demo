package main

import "fmt"

/*
	类型断言
*/

type Usber interface {
    start()
    stop()
}

// Computer 电脑
type Computer struct {
}

func (c Computer) work(usb Usber) {
    // 类型断言: 判断 usb 的类型
    if _, ok := usb.(Phone); ok {
        usb.start()
    } else {
        usb.stop()
    }
}

// Phone 手机
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
    computer := Computer{}

    phone := Phone{
        Name: "小米手机",
    }

    camera := Camera{}

    computer.work(phone)
    computer.work(camera)
}
