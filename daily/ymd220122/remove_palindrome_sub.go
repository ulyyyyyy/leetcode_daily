package ymd220122

func removePalindromeSub(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}
	for i, n := 0, length; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return 2
		}
	}
	return 1
}
