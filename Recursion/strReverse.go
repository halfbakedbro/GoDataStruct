package main

import (
	"fmt"
	"time"
)

func timetrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %dns\n", name, elapsed.Nanoseconds())
}

func reverseStringResursion(str string) string {
	// Without this, the program generates slice bound error
	if len(str) <= 1 {
		return str
	}
	return reverseStringResursion(string(str[1:])) + string(str[0])
}

func strRevItr(str string) {
	run := []rune(str)

	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {

		run[i], run[j] = run[j], run[i]
	}

	fmt.Println(string(run))
}

func main() {

	str := "This is string slice"
	fmt.Println(reverseStringResursion(str))

}
