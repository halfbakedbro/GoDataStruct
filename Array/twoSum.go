func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		_, prs := m[n]
		if prs {
			return []int{m[n], i}
		} else {
			m[target-n] = i
		}
	}
	return nil
}