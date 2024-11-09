package main

import "fmt"

/*
结构体嵌套
*/

type User struct {
	Username string
	Password string
	Address  Address // 嵌套结构体
}

type Address struct {
	Name  string
	Phone string
	City  string
}

type User2 struct {
	Username string
	Password string
	Address  // 嵌套 匿名结构体
}

func main() {
	var u User
	u.Username = "张三"
	u.Password = "123"

	u.Address.Name = "老六"
	u.Address.Phone = "987654321"
	u.Address.City = "北京"
	fmt.Printf("u=#%v\n", u)
	fmt.Println("- - - - - - - - - - - - - - - - - - - - ")

	var u2 User2
	u2.Username = "李四"
	u2.Password = "456"

	u2.Address.Name = "老七"
	u2.Address.Phone = "123456789"
	u2.City = "上海" // 当访问结构体成员时，优先访问当前结构体的成员，如果找不到，再去访问嵌套的"匿名结构体"的成员
	fmt.Printf("u2=#%v\n", u2)

}
