package main

import "fmt"

// 匿名字段
// 这么写，类型必须是唯一的, 否则报错
type Person0 struct {
	string
	int
	// string // 报错：string redeclared
}

/*
结构体的字段类型
 1. 结构体的字段类型可以是：基本数据类型、切片、Map、结构体
 2. 如果结构体字段是：指针、slice、map, 那么这些字段的零值是nil, 即还没有分配内存空间
    如果要使用这样的字段，必须先分配内存空间，然后才能使用
*/
type Person struct {
	Name  string
	Age   int
	Hobby []string
	map1  map[string]string
}

func main() {
	p0 := Person0{
		"张三",
		20,
	}
	fmt.Println(p0)
	fmt.Println("- - - - - - - - - - - - - - - - - - - - ")

	p := Person{
		Name: "张三",
		Age:  20,
	}
	p.Hobby = make([]string, 3, 6)
	p.Hobby[0] = "写代码"
	p.Hobby[1] = "打篮球"
	p.Hobby[2] = "睡觉"

	p.map1 = make(map[string]string)
	p.map1["address"] = "北京"
	p.map1["phone"] = "123"

	fmt.Printf("p=%#v\n", p)
	fmt.Println("- - - - - - - - - - - - - - - - - - - - ")
}
