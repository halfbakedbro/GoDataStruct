package recursive

import (
	"fmt"
	"time"
)

func timetrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %dns\n", name, elapsed.Nanoseconds())
}

func FactRecursive(num int64) int64 {
	//defer timetrack(time.Now(), "factorial recursive")
	if num == 1 {
		return 1
	}
	return num * FactRecursive(num-1)
}

func FactIterative(num int) int {

	defer timetrack(time.Now(), "factorial iterative")
	sum := 1

	for num != 1 {
		sum = num * sum
		num = num - 1
	}

	return sum
}

func main() {
	//fmt.Println("Hello, playground")
	//fmt.Println(FactIterative(9))
	now := time.Now()
	fmt.Println(FactRecursive(9))
	elapsed := time.Since(now)
	fmt.Printf("took %dns\n", elapsed.Nanoseconds())
}
