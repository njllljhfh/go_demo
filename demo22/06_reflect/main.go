package main

import (
    "fmt"
    "reflect"
)

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

func (s *Student) Print() {
    fmt.Println("这是一个打印方法...")
}

// PrintStructField 打印字段
func PrintStructField(s interface{}) {
    t := reflect.TypeOf(s) // reflect.TypeOf() 获取 类型变量
    fmt.Printf("t = %v\n", t)
    fmt.Printf("t.Kind() = %v\n", t.Kind())

    // 判断参数是不是结构体类型(支持 结构体对象 和 结构体指针)
    if (t.Kind() != reflect.Ptr && t.Kind() != reflect.Struct) || (t.Kind() == reflect.Ptr && t.Elem().Kind() != reflect.Struct) {
        fmt.Println("传入的参数不是一个结构体")
        return
    }

    if t.Kind() == reflect.Ptr {
        fmt.Printf("t.Elem() = %v\n", t.Elem()) // 当s是结构体变量而不是指针时，报错。panic: reflect: Elem of invalid type main.Student
        t = t.Elem()
    }
    fmt.Println("- - -")

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

    // v := reflect.ValueOf(s) // reflect.ValueOf() 获取 值变量
    // 通过 值变量 获取结构体属性对应的值
}

func main() {
    stu1 := Student{
        Name:  "小明",
        Age:   15,
        Score: 98,
    }

    // PrintStructField(stu1)
    PrintStructField(&stu1)

    // 测试代码的健壮性
    // a := 1
    // PrintStructField(a)
    // PrintStructField(&a)
}
