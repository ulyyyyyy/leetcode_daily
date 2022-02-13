package ymd220213

import "math"

// 给你一个字符串 text，你需要使用 text 中的字母来拼凑尽可能多的单词 "balloon"（气球）。
//
// 字符串 text 中的每个字母最多只能被使用一次。请你返回最多可以拼凑出多少个单词 "balloon"。

func maxNumberOfBalloons(text string) (ans int) {
	inBalloon := func(char rune) bool {
		return char == 'b' || char == 'a' || char == 'l' || char == 'o' || char == 'n'
	}

	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}

	tmpMap := map[rune]int{
		'b': 0, // 1
		'a': 0, // 1
		'l': 0, // 2
		'o': 0, // 2
		'n': 0, // 1
	}
	for _, r := range text {
		if inBalloon(r) {
			tmpMap[r]++
		}
	}
	ans = math.MaxInt
	for r, i := range tmpMap {
		if r == 'l' || r == 'o' {
			ans = min(i/2, ans)
		}
		ans = min(i, ans)
	}
	return
}
