package ymd220131

// 给你一个非负整数 num ，请你返回将它变成 0 所需要的步数。 如果当前数字是偶数，你需要把它除以 2 ；否则，减去 1 。

func numberOfSteps(num int) (ans int) {
	for ; num > 0; num >>= 1 {
		ans += num & 1
		if num > 1 {
			ans++
		}
	}
	return
}
