package main

import "fmt"

// 结构体，是值类型
type Person struct { // 首字母大写表示公有, 首字母小写表示私有
	name string
	age  int
	sex  string
}

func main() {
	// 实例化结构体 --- 方法1
	var p1 Person // 实例化结构体
	p1.name = "张三"
	p1.sex = "男"
	p1.age = 20
	fmt.Printf("p1的值:%v, p1的类型:%T\n", p1, p1)  //p1的值:{张三 20 男}, p1的类型:main.Person
	fmt.Printf("p1的值:%#v, p1的类型:%T\n", p1, p1) // %#v 打印详细信息
	fmt.Println("- - - - - - - - - - - - - - - - - - - - ")

	// 实例化结构体 --- 方法2
	var p2 = new(Person) // p2是指针
	p2.name = "李四"       // 底层是 (*p2).name = "李四"
	p2.age = 20
	p2.sex = "男"
	fmt.Printf("p2的值:%#v, p2的类型:%T\n", p2, p2) //p2的值:&main.Person{name:"李四", age:20, sex:"男"}, p2的类型:*main.Person
	fmt.Println("- - - - - - - - - - - - - - - - - - - - ")

	// 实例化结构体 --- 方法3
	var p3 = &Person{} // p3是指针
	p3.name = "王五"
	p3.age = 20
	p3.sex = "男"
	fmt.Printf("p3的值:%#v, p3的类型:%T\n", p2, p2)
	fmt.Println("- - - - - - - - - - - - - - - - - - - - ")

	// 实例化结构体 --- 方法4
	p4 := Person{
		name: "赵六",
		age:  20,
		sex:  "男",
	}
	fmt.Printf("p4的值:%#v, p4的类型:%T\n", p4, p4)
	fmt.Println("- - - - - - - - - - - - - - - - - - - - ")

	// 实例化结构体 --- 方法5
	p5 := &Person{
		name: "老七",
		age:  20,
		sex:  "男",
	} // p5是指针
	fmt.Printf("p5的值:%#v, p5的类型:%T\n", p5, p5)
	fmt.Println("- - - - - - - - - - - - - - - - - - - - ")

	// 实例化结构体 --- 方法6
	// 没有指定字段，默认值为该字段的 0值
	p6 := &Person{
		name: "老刘",
	} // p6是指针
	fmt.Printf("p6的值:%#v, p6的类型:%T\n", p6, p6)
	fmt.Println("- - - - - - - - - - - - - - - - - - - - ")

	// 实例化结构体 --- 方法7
	// 不指定key，则按顺序赋值
	p7 := &Person{
		"老六",
		20,
		"男",
	} // p7是指针
	fmt.Printf("p7的值:%#v, p7的类型:%T\n", p7, p7)
	fmt.Println("- - - - - - - - - - - - - - - - - - - - ")
}
