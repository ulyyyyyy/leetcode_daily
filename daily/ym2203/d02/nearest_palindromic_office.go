package d02

import (
	"math"
	"strconv"
)

func nearestPalindromic(n string) string {
	m := len(n)
	// 找到上一位 & 下一位的第一个回文数
	candidates := []int{int(math.Pow10(m-1)) - 1, int(math.Pow10(m)) + 1}
	// 获取前缀
	selfPrefix, _ := strconv.Atoi(n[:(m+1)/2])
	// 遍历前缀左右的两位数，找到前缀能够构成的回文数
	for _, x := range []int{selfPrefix - 1, selfPrefix, selfPrefix + 1} {
		y := x
		if m&1 == 1 {
			y /= 10
		}
		for ; y > 0; y /= 10 {
			x = x*10 + y%10
		}
		candidates = append(candidates, x)
	}

	ans := -1
	selfNumber, _ := strconv.Atoi(n)
	for _, candidate := range candidates {
		if candidate != selfNumber {
			// 判断是否是最近的值
			if ans == -1 || abs(candidate-selfNumber) < abs(ans-selfNumber) || abs(candidate-selfNumber) == abs(ans-selfNumber) && candidate < ans {
				ans = candidate
			}
		}
	}
	return strconv.Itoa(ans)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
