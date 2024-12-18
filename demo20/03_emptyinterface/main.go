package main

import "fmt"

// A 空接口，可以表示任何类型
type A interface{}

// 空接口作为函数参数
func show(a interface{}) {
    fmt.Printf("值:%v 类型:%T\n", a, a)
}

func main() {
    var a A

    var str = "hello"
    a = str
    fmt.Printf("值:%v 类型:%T\n", a, a)

    var num = 20
    a = num
    fmt.Printf("值:%v 类型:%T\n", a, a)

    var flag = true
    a = flag
    fmt.Printf("值:%v 类型:%T\n", a, a)
    fmt.Println("-------------------------------")

    // 空接口也可以直接当做类型使用，可以接收任意类型的值
    var a2 interface{}
    a2 = 20
    fmt.Printf("值:%v 类型:%T\n", a2, a2)
    fmt.Println("-------------------------------")

    show(20)
    show("你好！")
    show([]int{1, 2, 3, 4, 5})
    fmt.Println("-------------------------------")

    // 空接口作为map的值
    var m1 = make(map[string]interface{})
    m1["username"] = "张三"
    m1["age"] = 20
    m1["married"] = true
    fmt.Println(m1)
    fmt.Println("-------------------------------")

    // 空接口作为切片的值
    var s1 = []interface{}{1, 2, "hello", true}
    fmt.Println(s1)
}
