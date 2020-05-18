func swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}

func reverse(nums *[]int, a int, b int) {

	for a < b {
		//swap(&nums[a], &nums[b])
		t := (*nums)[a]
		(*nums)[a] = (*nums)[b]
		(*nums)[b] = t
		a++
		b--
	}
}

func rotate(nums []int, k int) {
	k = k % len(nums)

	reverse(&nums, 0, len(nums)-1)
	reverse(&nums, k, len(nums)-1)
	reverse(&nums, 0, k-1)
}