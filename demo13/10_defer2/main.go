package main

import "fmt"

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x // y=5
}

func f4() (x int) {
	defer func(x int) {
		fmt.Printf("f4 defer 内 x=%d\n", x) // x=0
		x++
	}(x) // defer 注册要延迟执行的函数时，该函数的所有参数都需要确定其值
	return 5
}

func main() {
	fmt.Println(f1()) //5
	fmt.Println(f2()) //6
	fmt.Println(f3()) //5
	fmt.Println(f4()) //5
}