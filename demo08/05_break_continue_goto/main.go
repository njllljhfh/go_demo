package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		if i == 2 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("------------------------------------------")

	// 在多重循环中，可以用 label 给某一个循环命名，然后用 break 语句配合标签名，来指定要终止的是哪一个循环
label1:
	for i := 0; i < 2; i++ {
		for j := 0; j < 10; j++ {
			if j == 3 {
				break label1
			}
			fmt.Printf("i=%v, j=%v\n", i, j)
		}
	}
	fmt.Println("------------------------------------------")

	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("------------------------------------------")

	// 在 continue 语句后添加标签时，表示开始标签对应的循环
label2:
	for i := 0; i < 2; i++ {
		for j := 0; j < 10; j++ {
			if j == 3 {
				continue label2
			}
			fmt.Printf("i=%v, j=%v\n", i, j)
		}
	}
	fmt.Println("------------------------------------------")

	// goto 语句通过标签进行代码间的无条件跳转
	var n = 30
	if n > 24 {
		fmt.Printf("成年人\n")
		goto label3
	}
	fmt.Println("aaa")
	fmt.Println("bbb")
label3:
	fmt.Println("ccc")
	fmt.Println("ddd")
}
