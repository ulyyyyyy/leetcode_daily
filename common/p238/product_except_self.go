package p238

// productExceptSelf 暴力，保存所有非当前序号的数组，累乘
func productExceptSelf(nums []int) []int {
	l := len(nums)
	indexSli := make([][]int, l)
	retSli := make([]int, l)
	for i := range nums {
		indexSli[i] = append(append([]int{}, nums[:i]...), append([]int{}, nums[i+1:]...)...)
	}

	for i, s := range indexSli {
		t := 1
		for _, n := range s {
			t *= n
		}
		retSli[i] = t
	}
	return retSli
}
