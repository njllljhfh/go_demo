package something

import "fmt"

func Factorial(x int) (res int) {
	res = 1
	for i := 1; i <= x; i++ {
		res *= i
	}
	return
}

func init() {
	fmt.Println("something init...")
}
