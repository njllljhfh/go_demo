package main

import "fmt"

/*
结构体嵌套
*/
type User struct {
	Username string
	Password string
	AddTime  string // 重名字段
	Address         // 嵌套 匿名结构体
}

type Address struct {
	Name    string
	Phone   string
	City    string
	AddTime string // 重名字段
}

func main() {
	var u User
	u.Username = "tying"
	u.Password = "1234567"

	u.Address.Name = "张三"
	u.Address.Phone = "987654321"
	u.Address.City = "北京"

	u.AddTime = "2020-05-1" // 改的是User下的AddTime字段
	u.Address.AddTime = "2020-06-1"
	fmt.Printf("u=%#v\n", u)
	fmt.Println("- - - - - - - - - - - - - - - - - - - - ")
}
