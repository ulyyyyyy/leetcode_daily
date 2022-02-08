package ymd220201

import "unicode"

// 当一个字符串 s 包含的每一种字母的大写和小写形式 同时 出现在 s 中，就称这个字符串 s 是 美好 字符串。比方说，"abABB" 是美好字符串，因为 'A' 和 'a' 同时出现了，且 'B' 和 'b' 也同时出现了。然而，"abA" 不是美好字符串因为 'b' 出现了，而 'B' 没有出现。
//
// 给你一个字符串 s ，请你返回 s 最长的 美好子字符串 。如果有多个答案，请你返回 最早 出现的一个。如果不存在美好子字符串，请你返回一个空字符串。

func longestNiceSubstring(s string) (ans string) {
	if s == "" {
		return
	}
	lower, upper := 0, 0
	for _, ch := range s {
		if unicode.IsLower(ch) {
			lower |= 1 << (ch - 'a')
		} else {
			upper |= 1 << (ch - 'A')
		}
	}
	if lower == upper {
		return s
	}
	valid := lower & upper
	for i := 0; i < len(s); i++ {
		start := i
		for i < len(s) && valid>>(unicode.ToLower(rune(s[i]))-'a')&1 == 1 {
			i++
		}
		if t := longestNiceSubstring(s[start:i]); len(t) > len(ans) {
			ans = t
		}
	}
	return
}
