package main

import "fmt"

/*
结构体嵌套
*/
type User struct {
	Username string
	Password string
	Address
	Email
}

type Address struct {
	Name    string
	Phone   string
	City    string
	AddTime string // 重名字段
}

type Email struct {
	Account string
	AddTime string // 重名字段
}

func main() {
	var u User
	u.Username = "tying"
	u.Password = "1234567"

	u.Address.Name = "张三"
	u.Address.Phone = "987654321"
	u.Address.City = "北京"

	// 当访问结构体成员时，优先访问当前结构体的成员，如果找不到，再去访问嵌套的"匿名结构体"的成员
	// u.AddTime = "2020-05-1" // ambiguous selector u.AddTime

	u.Address.AddTime = "2020-05-1"
	u.Email.AddTime = "2020-06-1"
	fmt.Printf("u=%#v\n", u)
	fmt.Println("- - - - - - - - - - - - - - - - - - - - ")
}
