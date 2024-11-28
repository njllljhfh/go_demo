package main

import (
	"fmt"
	"reflect"
)

/*
   反射：获取结构体字段，结构体方法，调用结构体方法
*/

type Student struct {
	Name  string `json:"name1" form:"username"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func (s *Student) GetInfo() string {
	str := fmt.Sprintf("姓名:%v, 年龄:%v, 成绩:%v\n", s.Name, s.Age, s.Score)
	return str
}

func (s *Student) SetInfo(name string, age int, score int) {
	s.Name = name
	s.Age = age
	s.Score = score
}

func (s Student) Print() {
	fmt.Println("这是一个打印方法...")
}

// PrintStructField 打印字段
func PrintStructField(s interface{}) {
	t := reflect.TypeOf(s) // reflect.TypeOf() 获取 类型变量
	fmt.Printf("t = %v\n", t)
	fmt.Printf("t.Kind() = %v\n", t.Kind())

	v := reflect.ValueOf(s) // reflect.ValueOf() 获取 值变量
	fmt.Printf("v.Kind() = %v\n", v.Kind())
	fmt.Println("- - -")

	// 判断参数是不是结构体类型(支持 结构体对象 和 结构体指针)
	if (t.Kind() != reflect.Ptr && t.Kind() != reflect.Struct) || (t.Kind() == reflect.Ptr && t.Elem().Kind() != reflect.Struct) {
		fmt.Println("传入的参数不是一个结构体")
		fmt.Println("------------------------")
		return
	}

	if t.Kind() == reflect.Ptr {
		fmt.Printf("t.Elem() = %v\n", t.Elem()) // t.Elem() 当s是结构体变量而不是指针时，报错。panic: reflect: Elem of invalid type main.Student
		t = t.Elem()

		fmt.Printf("v.Elem() = %v\n", v.Elem()) // v.Elem() 当s是结构体变量而不是指针时，报错。panic: reflect: call of reflect.Value.Elem on struct Value
		fmt.Printf("v.Elem().Kind() = %v\n", v.Elem().Kind())
		v = v.Elem()
	}
	fmt.Println("- - - - - - - - - - - - - - - - - - ")

	// 通过类型变量里面的 Field 可以获取结构体的字段
	field0 := t.Field(0) // 0：表示获取第0个属性
	fmt.Printf("%#v\n", field0)
	fmt.Println("字段名称：", field0.Name)
	fmt.Println("字段类型：", field0.Type)
	fmt.Println("字段Tag：", field0.Tag)
	fmt.Println("字段Tag：", field0.Tag.Get("json"))
	fmt.Println("字段Tag：", field0.Tag.Get("form"))
	fmt.Println("- - -")

	// 通过类型变量里面的 FieldByName 可以获取结构体的字段
	field1, ok := t.FieldByName("Age") // Age 是结构体中定义的属性名称
	if ok {
		fmt.Println("字段名称：", field1.Name)
		fmt.Println("字段类型：", field1.Type)
		fmt.Println("字段Tag：", field1.Tag.Get("json"))
	}
	fmt.Println("- - -")

	// 通过类型变量里面的 NumField 可以获取该结构体有几个字段
	fieldCount := t.NumField()
	fmt.Printf("结构体有 %v 个属性\n", fieldCount)
	fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - -")

	// 通过 值变量 获取结构体属性对应的值
	fmt.Println(v.FieldByName("Name"))
	fmt.Println(v.FieldByName("Age"))
	fmt.Println("- - -")
	fmt.Println("通过反射，遍历获取结构体的字段信息")
	for i := 0; i < fieldCount; i++ {
		fmt.Printf("属性名称:%v, 属性值:%v, 属性类型:%v, 属性Tag:%v\n", t.Field(i).Name, v.Field(i), t.Field(i).Type, t.Field(i).Tag.Get("json"))
	}
}

// PrintStructFn 打印执行方法
func PrintStructFn(s interface{}) {
	t := reflect.TypeOf(s)  // reflect.TypeOf() 获取 类型变量
	v := reflect.ValueOf(s) // reflect.ValueOf() 获取 值变量

	// 判断参数是不是结构体类型(支持 结构体对象 和 结构体指针)
	if (t.Kind() != reflect.Ptr && t.Kind() != reflect.Struct) || (t.Kind() == reflect.Ptr && t.Elem().Kind() != reflect.Struct) {
		fmt.Println("传入的参数不是一个结构体")
		fmt.Println("------------------------")
		return
	}

	// 通过 类型变量 的 Method 可以获取结构体方法
	method0 := t.Method(0)    // 参数 0：方法的顺序 跟 方法在代码中的编写顺序无关，而是与 方法名的 ASCII码 有关。
	fmt.Println(method0.Name) // GetInfo
	fmt.Println(method0.Type) // func(*main.Student) string
	fmt.Println("- - -")

	// 通过 类型变量 获取这个结构体有多少个方法
	method1, ok := t.MethodByName("Print")
	if ok {
		fmt.Println(method1.Name) // Print
		fmt.Println(method1.Type) // func(*main.Student)

	}
	fmt.Println("- - -")

	// 通过 值变量 执行方法（注意：需要使用 值变量，并且要注意参数）v.Method(0).Call(nil) 或 v.MethodByName("Print").Call(nil)
	v.Method(1).Call(nil) // nil 表示没有参数
	v.MethodByName("Print").Call(nil)
	info := v.MethodByName("GetInfo").Call(nil) // Call() 返回值的类型是 []reflect.Value 切片
	fmt.Println(info)
	fmt.Println("- - -")

	// 执行方法传入参数（注意：需要使用 值变量，并且要注意参数，接收的参数是 []reflect.Value 类型的切片）
	var params []reflect.Value
	params = append(params, reflect.ValueOf("李四"))
	params = append(params, reflect.ValueOf(23))
	params = append(params, reflect.ValueOf(99))
	v.MethodByName("SetInfo").Call(params) // 执行 SetInfo 方法，传入参数
	info2 := v.MethodByName("GetInfo").Call(nil)
	fmt.Println(info2)
	fmt.Printf("len(info2) = %v\n", len(info2))
	fmt.Println("- - -")

	// 获取方法数量
	fmt.Printf("方法的数量: %v\n", t.NumMethod())
	fmt.Printf("方法的数量: %v\n", v.NumMethod())
	fmt.Println("- - -")
}

func main() {
	stu1 := Student{
		Name:  "小明",
		Age:   15,
		Score: 98,
	}

	// PrintStructField(stu1)
	PrintStructField(&stu1)

	// 测试 PrintStructField 的健壮性
	// a := 1
	// PrintStructField(a)
	// PrintStructField(&a)
	fmt.Println("==============================================================================")

	// PrintStructFn(stu1)
	PrintStructFn(&stu1) // 函数内部有修改结构体的逻辑，所以要传 结构体的指针变量

	// stu1 已被修改
	fmt.Printf("%#v\n", stu1) // main.Student{Name:"李四", Age:23, Score:99}
}
