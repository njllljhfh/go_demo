package tools

import (
	"fmt"
	"itying/something"
)

func Mul(x, y int) int {
	return x * y
}

func init() {
	fmt.Println("tools init...")
}

func PrintData() {
	res := something.Factorial(5)
	fmt.Printf("5的阶乘：%d\n", res)
}
