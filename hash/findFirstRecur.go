package main

import "fmt"

func firstRecur(nums []int) (int, bool) {
	mar := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		if !mar[nums[i]] {
			mar[nums[i]] = true
		} else {
			return nums[i], true
		}
	}

	return 0, false

}

func main() {
	arr := []int{2, 1, 4, 5, 2, 9, 7, 11}

	in, ok := firstRecur(arr)
	if ok {
		fmt.Println(in)
	} else {
		fmt.Println("no recurring")
	}
}
