package d04

// subArrayRanges 直接双指针遍历，记录每个状态下的最值，直接相减后累加
func subArrayRanges(nums []int) (ans int64) {
	size := len(nums)
	for i := 0; i < size; i++ {
		max, min := nums[i], nums[i]
		for j := i + 1; j < size; j++ {
			max = maxFunc(max, nums[j])
			min = minFunc(min, nums[j])
			ans += int64(max - min)
		}
	}
	return
}

func maxFunc(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func minFunc(a, b int) int {
	if a >= b {
		return b
	}
	return a
}
