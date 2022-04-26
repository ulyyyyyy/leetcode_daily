package p31

func nextPermutation(nums []int) {
	l := len(nums)

	leftIdx := -1
	// 找到左侧需要交换的元素
	for i := l - 1; i > 0; i-- {
		j := i - 1
		if nums[j] < nums[i] {
			leftIdx = j
			break
		}
	}

	if leftIdx == -1 {
		reverse(nums)
		return
	}
	// 找到
	for i := l - 1; i > 0; i-- {
		if nums[i] > nums[leftIdx] {
			nums[i], nums[leftIdx] = nums[leftIdx], nums[i]
			break
		}
	}
	reverse(nums[leftIdx+1:])
}

func reverse(nums []int) {
	l := len(nums)
	for i := 0; i < l/2; i++ {
		nums[i], nums[l-1-i] = nums[l-1-i], nums[i]
	}
}
