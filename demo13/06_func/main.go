package main

import "fmt"

//函数作为返回值

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

type caleType func(int, int) int

func do(o string) caleType {
	switch o {
	case "+":
		return add
	case "-":
		return sub
	case "*":
		// 匿名函数
		return func(x, y int) int {
			return x * y
		}
	default:
		return nil
	}
}

func main() {
	a := do("+")
	fmt.Println(a(12, 4))

	b := do("*")
	fmt.Println(b(3, 4))

	// 匿名函数, 直接调用
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)
}
