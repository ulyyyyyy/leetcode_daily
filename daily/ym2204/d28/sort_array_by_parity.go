package d28

func sortArrayByParity(nums []int) []int {
	l := len(nums)
outer:
	for i := 0; i < l; i++ {
		if nums[i]%2 == 0 {
			continue
		}
		for j := l - 1; j > i; j-- {
			if nums[j]%2 == 0 {
				nums[i], nums[j] = nums[j], nums[i]
				continue outer
			}
		}
		return nums
	}
	return nums
}
