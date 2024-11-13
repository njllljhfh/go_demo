package main

import (
	"fmt"

	"github.com/shopspring/decimal"
	"github.com/tidwall/gjson"
)

// 用第三方包
// pkg.go.dev  // 该网站可以搜索go的第三方包

func main() {
	fmt.Println("Hello World!")
	price, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}
	fmt.Println(price)
	fmt.Println("- - - - - - - - - - - - - - - - - - -")

	const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
	value := gjson.Get(json, "name.last")
	println(value.String())
}

/*
1. go mod int 项目名      // 初始化项目
2. go mod tidy           // 下载依赖
*/
