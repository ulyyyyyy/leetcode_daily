package d19

func shortestToChar(s string, c byte) []int {
	n := len(s)
	ans := make([]int, n)

	// 从左边遍历，保存上一个目标字符位置
	idx := -n
	for i, ch := range s {
		if byte(ch) == c {
			idx = i
		}
		ans[i] = i - idx
	}

	// 从右边倒序遍历，保存上一个目标字符位置
	idx = n * 2
	for i := n - 1; i >= 0; i-- {
		if s[i] == c {
			idx = i
		}
		ans[i] = min(ans[i], idx-i)
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
