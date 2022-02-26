package d26

func maximumDifference(nums []int) (max int) {
	max = -1
	last := len(nums) - 1

	for i := len(nums) - 2; i > -1; i-- {
		if nums[i] < last {
			max = maxFunc(max, last-nums[i])
		}
		last = maxFunc(last, nums[i])
	}
	return
}

func maxFunc(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
