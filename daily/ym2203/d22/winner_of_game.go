package d22

var (
	orderList []rune
)

// winnerOfGame 直接模拟，会超时
func winnerOfGame(colors string) bool {
	orderList = []rune{'A', 'B'}

outer:
	for {
		curUser := orderList[0]
		if len(colors) <= 2 {
			return curUser != 'A'
		}
		orderList[0], orderList[1] = orderList[1], orderList[0]
		for i := 1; i < len(colors)-1; i++ {
			if rune(colors[i]) == curUser && rune(colors[i-1]) == curUser && rune(colors[i+1]) == curUser {
				// 则可以删除
				colors = colors[:i] + colors[i+1:]
				continue outer
			}
		}
		return curUser != 'A'
	}
}

// 滑动窗口 找到A, B最多相邻N个元素的
func winnerOfGame2(colors string) bool {
	if len(colors) <= 2 {
		return false
	}
	total := [2]int{0, 0}
	// 记录当前元素与连续元素个数
	cur, cnt := '0', 0
	for _, r := range colors {
		// 如果当前元素切换，那么重新计数
		if cur != r {
			cur, cnt = r, 1
		} else {
			cnt++
			// 当连续个数超过三个时候，总数++
			if cnt >= 3 {
				total[cur-'A']++
			}
		}
	}
	return total[0] > total[1]
}
