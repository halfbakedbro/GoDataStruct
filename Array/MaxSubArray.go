const UintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64

const (
	MaxInt  = 1<<(UintSize-1) - 1 // 1<<31 - 1 or 1<<63 - 1
	MinInt  = -MaxInt - 1         // -1 << 31 or -1 << 63
	MaxUint = 1<<UintSize - 1     // 1<<32 - 1 or 1<<64 - 1
)

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxSubArray(nums []int) int {

	curSum, best := 0, MinInt

	for i := 0; i < len(nums); i++ {
		curSum = Max(curSum+nums[i], nums[i])
		best = Max(curSum, best)
	}

	return best
}