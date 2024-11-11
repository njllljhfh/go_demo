package calc

// 自定义的包

// 首字母大写，才能被外部调用
var a = 100  // 私有变量
var Age = 18 // 公有变量

// 公有方法
func Add(x, y int) int {
	return x + y
}

// 公有方法
func Sub(x, y int) int {
	return x - y
}
