package Sorting

func SelectionSort(arr []int) []int {

	smallest := 0

	for i := 0; i < len(arr); i = i + 1 {
		smallest = i

		for j := i + 1; j < len(arr); j = j + 1 {
			if arr[j] < arr[smallest] {
				smallest = j
			}
		}

		arr[i], arr[smallest] = arr[smallest], arr[i]
	}

	return arr
}
