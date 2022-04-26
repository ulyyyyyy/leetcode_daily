package d22

func maxRotateFunction(nums []int) int {
	// F(k) = F(k-1) + numSum - n * nums[n-k]
	max := 0
	f := 0
	numSum := 0
	for i, n := range nums {
		f += i * n
		numSum += n
	}

	n := len(nums)
	for i := 0; i < n; i++ {
		fTmp := f + numSum - n*nums[n-i]
		f = fTmp
		max = maxFunc(max, fTmp)
	}
	return max
}

func maxFunc(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
