package d25

import "math"

// 给定一个整数 n ，返回 n! 结果中尾随零的数量。
//
// 提示 n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1

// trailingZeroes 直接模拟，遍历所有的字符
func trailingZeroes(n int) (cnt int) {
	for i := 5; i <= n; i++ {
		// 判断 i 是否包含 5
		for j := float64(i); j > 0; j /= 5 {
			if math.Floor(j) == j && int(j)%5 == 0 {
				cnt++
			}
		}
	}
	return
}

// trailingZeroes2 上一个方法因为遍历了许多无关的参数，所有耗时增加。时间复杂度 O(n/5)
// 优化，只遍历 5 个倍数个
func trailingZeroes2(n int) (cnt int) {
	// 每 5 个遍历一次。
	for i := 5; i <= n; i += 5 {
		// 判断该值有多少个 5
		for j := i; j%5 == 0; j /= 5 {
			cnt++
		}
	}
	return
}

// trailingZeroes3 实现对数时间复杂度的算法
func trailingZeroes3(n int) (cnt int) {
	for n > 0 {
		n /= 5
		cnt += n
	}
	return
}
