package d10

import (
	"fmt"
)

// 给你一个整数 n ，请你返回所有 0 到 1 之间（不包括 0 和 1）满足分母小于等于  n 的 最简 分数 。分数可以以 任意 顺序返回。
func simplifiedFractions(n int) (rlts []string) {
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			if getMax(i, j) {
				continue
			}
			rlts = append(rlts, fmt.Sprintf("%d/%d", j, i))

		}
	}
	return
}

// 判断是否有最大公倍数
func getMax(m, n int) bool {
	for {
		r := m % n
		if r == 0 {
			break
		}
		m, n = n, m%n
	}
	return n != 1
}
