package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1 string = "你好golang1"
	var str2 = "你好golang2"
	str3 := "你好golang3"

	fmt.Printf("%v--%T\n", str1, str1)
	fmt.Printf("%v--%T\n", str2, str2)
	fmt.Printf("%v--%T\n", str3, str3)
	fmt.Println("------------------------------------------")

	// 输出反斜杠
	str4 := "C:\\Go\\bin"
	fmt.Println(str4)
	fmt.Println("------------------------------------------")

	str5 := `
	多行字符串
	多行字符串
	多行字符串
	`
	fmt.Println(str5)
	fmt.Println("------------------------------------------")

	str6 := "你好"
	fmt.Println(len(str6)) // 6

	str7 := "aaaa"
	fmt.Println(len(str7)) // 4
	fmt.Println("------------------------------------------")

	// + 或 fmt.Sprintf 拼接字符串
	str8 := "你好"
	str9 := "golong"
	fmt.Println(str8 + str9)

	str10 := fmt.Sprintf("%s %s", str8, str9)
	fmt.Println(str10)

	// 长字符串的拼接
	str11 := "123" +
		"456" +
		"789"
	fmt.Println(str11)
	fmt.Println("------------------------------------------")

	// 分割字符串
	str12 := "123-456-789"
	arr := strings.Split(str12, "-") // arr 是切片
	fmt.Println(arr)
	fmt.Println("------------------------------------------")

	// 把切片连接成字符串
	arr2 := []string{"php", "java", "golang"}
	str13 := strings.Join(arr2, "-")
	fmt.Println(str13)
	fmt.Println("------------------------------------------")

	// 判断字符串是否包含指定的字符串
	str14 := "this is str"
	str15 := "this"
	flag := strings.Contains(str14, str15) //str14 是否包含 str15
	fmt.Println(flag)
	fmt.Println("------------------------------------------")

	// 判断前缀后缀
	str16 := "this is str"
	str17 := "this"
	flag = strings.HasPrefix(str16, str17)
	fmt.Println(flag)
	flag = strings.HasSuffix(str16, str17)
	fmt.Println(flag)
	fmt.Println("------------------------------------------")

	str18 := "this is str"
	str19 := "is"
	num := strings.Index(str18, str19)
	fmt.Println(num)
	num = strings.LastIndex(str18, str19)
	fmt.Println(num)
	num = strings.LastIndex(str18, "666") // 查找不到返回 -1
	fmt.Println(num)
	fmt.Println("------------------------------------------")

	str20 := "你好啊"
	num20 := strings.Index(str20, "啊") // 6
	fmt.Println(num20)
}
