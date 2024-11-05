package main

import "fmt"

// defer

// defer 在命名返回值函数 和 匿名返回值函数 中的表现不一样

func f1(){
	fmt.Println("开始")

	defer func(){
		fmt.Println("aaaa")
		fmt.Println("bbbb")
	}()
	fmt.Println("结束")

}

// 匿名返回值函数
func f2() int {
	var a int
	defer func(){
		a++
	}()
	return a  // 0
}

// 命名返回值函数
func f3() (a int) {
	defer func(){
		a++
	}()
	return a // 1
}

/*
	在 Go 语言中，return x 语句分为两步执行。
	当一个函数执行到 return 语句时，它会首先将返回值（在这个例子中是 x的值）赋值给函数的命名返回值（如果有的话），
	然后执行 defer 语句（如果有的话），最后函数才会真正返回。

	如果函数是匿名返回，即没有显式地指定返回值的名称，那么 return 语句会直接返回函数体中计算出的值。
	在这种情况下，defer 语句不会影响返回值，因为没有命名返回值可以被修改。
*/

func main() {
	f1()
	fmt.Println("------------------------------------------")	

	fmt.Println(f2())	
	fmt.Println("------------------------------------------")	

	fmt.Println(f3())	
	fmt.Println("------------------------------------------")	
}