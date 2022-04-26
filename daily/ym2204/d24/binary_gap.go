package d24

func binaryGap(n int) int {
	pre := -1

	idx := 0
	max := 0
	for n >= 2 {
		t := n % 2
		// 如果当前数为1，那么开始找
		if t == 1 {
			// 如果存在上一个数字，那么维护距离
			if pre != -1 {
				max = maxFunc(max, idx-pre)
			}
			pre = idx
		}
		n /= 2
		idx++
	}
	// 最后判断第一位
	if pre != -1 {
		max = maxFunc(max, idx-pre)
	}
	return max
}

func maxFunc(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
