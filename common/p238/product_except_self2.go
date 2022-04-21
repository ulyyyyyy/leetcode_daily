package p238

// productExceptSelf2 改进方法。维护两个积数组，一个前缀积，一个后缀积。
// 两个数组用以表示某位置的前缀积与后缀积（不包括该位置）
// 所以，除自身以外数组的乘积可以表示为 某位置前缀和 l 与某位置后缀和 r 的乘积 -> l * r
func productExceptSelf2(nums []int) (answer []int) {
	l := len(nums)

	left, right := make([]int, l), make([]int, l)
	answer = make([]int, l)

	// 前缀积
	lNum := 1
	for i, n := range nums {
		left[i] = lNum
		lNum *= n
	}

	rNum := 1
	for i := l - 1; i >= 0; i-- {
		right[i] = rNum
		rNum *= nums[i]
	}

	for i := range nums {
		answer[i] = left[i] * right[i]
	}
	return
}
