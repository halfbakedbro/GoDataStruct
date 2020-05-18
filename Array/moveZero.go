func swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}

func moveZeroes(nums []int) {
	k := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			swap(&nums[i], &nums[k])
			k++
		}
	}
}

func containsDuplicate(nums []int) bool {
	b := make(map[int]bool, len(nums))

	for _, v := range nums {
		if _, ok := b[v]; ok {
			return true
		}
		b[v] = true
	}

	return false
}