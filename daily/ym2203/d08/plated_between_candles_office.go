package d08

func platesBetweenCandlesOffice(s string, queries [][]int) []int {
	n := len(s)
	preSum := make([]int, n)
	left := make([]int, n)
	// sum 为统计 * 号数量，l 为从左往右数，右侧的 | 坐标位置
	sum, l := 0, -1
	for i, ch := range s {
		if ch == '*' {
			sum++
		} else {
			l = i
		}
		preSum[i] = sum
		left[i] = l
	}

	right := make([]int, n)
	// r 为从右往左数，右侧的 | 坐标位置
	for i, r := n-1, -1; i >= 0; i-- {
		if s[i] == '|' {
			r = i
		}
		right[i] = r
	}

	// 使用 right 和 left 维护盘子的右边界与左边界的值
	ans := make([]int, len(queries))
	for i, q := range queries {
		// 取到盘子对应的右边界与左边界的值，即限定盘子实际范围
		x, y := right[q[0]], left[q[1]]
		if x >= 0 && y >= 0 && x < y {
			// 直接用预处理的值统计
			ans[i] = preSum[y] - preSum[x]
		}
	}
	return ans
}
