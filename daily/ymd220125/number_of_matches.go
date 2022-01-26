package ymd220125

func numberOfMatches(n int) int {
	if n <= 2 {
		return n - 1
	}

	ans := 0
	for n != 1 {
		ans += n / 2
		n = (n + 1) / 2
	}
	return ans
}

func numberOfMatches2(n int) int {
	return n - 1
}
