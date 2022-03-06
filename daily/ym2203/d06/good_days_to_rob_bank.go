package d06

import "sort"

// 你和一群强盗准备打劫银行。给你一个下标从 0 开始的整数数组 security ，其中 security[i] 是第 i 天执勤警卫的数量。日子从 0 开始编号。同时给你一个整数 time 。
//
// 如果第 i 天满足以下所有条件，我们称它为一个适合打劫银行的日子：
//
// 第 i 天前和后都分别至少有 time 天。
// 第 i 天前连续 time 天警卫数目都是非递增的。
// 第 i 天后连续 time 天警卫数目都是非递减的。
// 更正式的，第 i 天是一个合适打劫银行的日子当且仅当：security[i - time] >= security[i - time + 1] >= ... >= security[i] <= ... <= security[i + time - 1] <= security[i + time].
//
// 请你返回一个数组，包含 所有 适合打劫银行的日子（下标从 0 开始）。返回的日子可以 任意 顺序排列。

func goodDaysToRobBank(security []int, time int) []int {
	ans := make([]int, 0)
	n := len(security)
	left := make([]int, n)
	right := make([]int, n)
	for i := 1; i < n; i++ {
		if security[i] <= security[i-1] {
			left[i] = left[i-1] + 1
		}
		if security[n-i-1] <= security[n-i] {
			right[n-i-1] = right[n-i] + 1
		}
	}

	for i := time; i < n-time; i++ {
		if left[i] >= time && right[i] >= time {
			ans = append(ans, i)
		}
	}
	return ans
}

// 直接模拟。时间复杂度 O(n^2)，空间复杂度 O(1)
func goodDaysToRobBank2(security []int, time int) (goodDays []int) {
	goodDays = make([]int, 0)
	s := len(security)
	for i := time; i < s-time; i++ {
		preSli := security[i-time : i+1]
		lastSli := security[i : i+time+1]

		if sort.SliceIsSorted(preSli, func(last, pre int) bool {
			return preSli[last] > preSli[pre]
		}) && sort.SliceIsSorted(lastSli, func(last, pre int) bool {
			return lastSli[last] < lastSli[pre]
		}) {
			goodDays = append(goodDays, i)
		}
	}
	return goodDays
}
