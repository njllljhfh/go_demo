package main

import "fmt"

/*
	空接口 和 类型断言 的使用细节

	本节内容，实际项目中，用的比较多
*/

type Address struct {
    Name  string
    Phone int
}

func main() {
    var userinfo = make(map[string]interface{})
    userinfo["name"] = "张三"
    userinfo["age"] = 18
    userinfo["hobby"] = []string{"睡觉", "吃饭"}

    fmt.Println(userinfo["age"])
    fmt.Println(userinfo["hobby"])
    fmt.Printf("%T\n", userinfo["hobby"])

    // 这里不能直接获取 hobby 的值
    // 报错：invalid operation: cannot index userinfo["hobby"] (map index expression of type interface{})
    // fmt.Println(userinfo["hobby"][1])
    hobby, _ := userinfo["hobby"].([]string) // 类型断言
    fmt.Println(hobby[1])
    fmt.Println("-----------------------------")

    address := Address{
        Name:  "李四",
        Phone: 123412332,
    }
    userinfo["address"] = address
    fmt.Println(userinfo["address"])

    // 这里不能直接获取 address 的值
    // 报错：userinfo["address"].Name undefined (type interface{} has no field or method Name)
    // name := userinfo["address"].Name
    name := userinfo["address"].(Address).Name // 类型断言
    fmt.Println(name)
    fmt.Println("-----------------------------")

    // a := 1
    // // 报错：invalid operation: a (variable of type int) is not an interface
    // if v, ok := a.(int); ok {
    //     fmt.Println(v)
    // } else {
    //     fmt.Println("类型断言失败")
    // }
}
