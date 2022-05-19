package p3

func lengthOfLongestSubstring(s string) int {
	pre, n := -1, len(s)
	ans := 0

	tmpMap := make(map[byte]int)

	for i := 0; i < n; i++ {
		if i != 0 {
			delete(tmpMap, s[i-1])
		}

		for pre+1 < n && tmpMap[s[pre+1]] == 0 {
			// 不断地移动右指针
			tmpMap[s[pre+1]]++
			pre++
		}
		ans = maxFunc(ans, pre-i+1)
	}
	return ans
}

func maxFunc(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
