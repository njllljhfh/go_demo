package main

import "fmt"

// 函数定义

func sum(x int, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

// 接受可变参数
func sum2(x ...int) int {
	// x 是一个切片
	// fmt.Printf("%v--%T\n", x, x)
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

// 可变参数 要放在 形参列表的最后一个位置
// 具名返回值
func sum3(x int, y ...int) (sum int) {
	fmt.Printf("x=%v -- %T\n", x, x)
	fmt.Printf("y=%v -- %T\n", y, y)

	sum = x
	for _, v := range y {
		sum += v
	}
	return
}

// 返回多个值
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

func main() {
	res1 := sum(12, 3)
	fmt.Println(res1)

	res2 := sum(20, 2)
	fmt.Println(res2)

	res3 := sum2(12, 34, 45, 46)
	fmt.Println(res3)
	fmt.Println("------------------------------------------")

	res4 := sum3(100, 1, 2, 3, 4)
	fmt.Println(res4)
	fmt.Println("------------------------------------------")

	a, b := calc(10, 2)
	fmt.Println(a, b)
}
