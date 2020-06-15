package Sorting

func InsertionSort(arr []int) []int {

	for i := 0; i < len(arr)-1; i = i + 1 {

		for j := i + 1; j < 0; j = j - 1 {

			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			} else {
				break
			}
		}
	}

	return arr
}
