package main

import "fmt"

func main() {
	/*
		Here we create an array that will hold
		exactly 7 int. By default array is zero valued
	*/
	var a [7]int

	// Can set a value at an index using the array[index] = value
	a[4] = 10 // O(1)
	a[3] = 2
	a[5] = 3

	//Builtin len returns the length of array
	fmt.Println(len(a))

	//Declare and intialize array in one go
	b := [3]int{1, 2, 3}
	fmt.Println(b)
	// we can also have multi dimensional array
	var z [2][2]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			z[i][j] = i * j
		}
	}

	fmt.Println(z)
}
