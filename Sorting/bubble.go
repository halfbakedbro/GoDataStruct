package Sorting

func BubbleSort(array []int) []int {

	for i := 0; i < len(array)-1; i = i + 1 {

		for j := i + 1; j < len(array); j = j + 1 {

			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}

	}

	return array

}
