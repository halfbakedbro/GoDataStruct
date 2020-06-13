package recursive

import (
	"fmt"
)

func fibonacciIterative(num int) {
	a, b := 0, 1
	sum := a
	for i := 1; i <= num; i++ {
		//fmt.Println(sum)
		a, b = b, a+b
		sum = a
	}
	fmt.Println(sum)
}

func fibonacciRecursive(num int) int {

	if num == 0 {
		return 0
	} else if num == 1 {
		return 1
	} else {
		return fibonacciRecursive(num-1) + fibonacciRecursive(num-2)
	}

}

func main() {
	fmt.Println("Hello World")
	fmt.Println(fibonacciRecursive(8))
}
