package tools

func Sort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// 同一个包不能用同名的函数
// func PringInfo() {
// 	fmt.Println("打印信息")
// }
