package main

import "fmt"

// 指针也是一个变量，它的值是另一个变量的地址。
// 指针变量也有自己的地址。

/*
	在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，否则我们的值就没办法存储。
	而对于值类型的声明不需要分配内存空间，是因为它们在声明的时候已经默认分配好了内存空间。
	要分配内存，就引出来今天的 new 和 make。
	Go语言中 new 和 make 是内建的两个函数，主要用来分配内存。

	可以用 make 来分配并且初始化类型为 slice、map、chan 的数据。
	new 用来分配值类型的内存空间，并且内存对应的值为类型零值，返回的是指向类型的指针。
*/

func fn1(x int) {
	x = 10
}

func fn2(x *int) {
	*x = 40
}

func main() {
	a := 10
	p := &a //p 是 *int 类型的指针变量
	fmt.Printf("a的值：%v, a的类型%T, a的地址%p\n", a, a, &a)
	fmt.Printf("p的值：%v, p的类型%T, p的地址%p\n", p, p, &p)
	fmt.Println("------------------------------------------")

	fmt.Println(*p) // 通过指针变量p访问a的值

	*p = 20 // 通过指针变量p修改变量a的值
	fmt.Printf("a = %v\n", a)
	fmt.Println("------------------------------------------")

	b := 5
	fn1(b)
	fmt.Println(b) // 5
	fn2(&b)
	fmt.Println(b) // 40
	fmt.Println("------------------------------------------")

	// 错误的写法
	// var userinfo map[string]string
	// userinfo["username"] = "张三"
	// fmt.Println(userinfo) // panic: assignment to entry in nil map

	// map 是 引用类型
	var userinfo = make(map[string]string)
	userinfo["username"] = "张三"
	fmt.Println(userinfo)
	fmt.Println("------------------------------------------")

	// 错误的写法
	// var s []int
	// s[0] = 1
	// fmt.Println(s) // panic: runtime error: index out of range [0] with length 0

	// slice 是 引用类型
	var s = make([]int, 5, 5)
	s[0] = 1
	fmt.Println(s)
	fmt.Println("------------------------------------------")

	// 错误的写法
	// var a1 *int // 这里只是声明了一个指针变量，并没有分配内存空间。
	// *a1 = 10
	// fmt.Println(a1) // panic: runtime error: invalid memory address or nil pointer dereference

	// 指针 是 引用类型
	// 开发中 new函数 不太常用，使用 new函数得到的是一个指针类型的指针变量。分配的内存中保存的值是该类型的0值
	var a1 = new(int) // 分配内存空间, a1 是 *int 类型的指针变量。a1指向的内存空间中保存的值是 0
	fmt.Printf("a1的值：%v, a1的类型：%T, a1指向的内存空间中保存的值：%v\n", a1, a1, *a1)

	*a1 = 10
	fmt.Println(a1)
	fmt.Println(*a1)
}
