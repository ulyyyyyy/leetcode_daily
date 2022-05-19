package d19

import (
	"sort"
)

// 给你一个长度为 n 的整数数组 nums ，返回使所有数组元素相等需要的最少移动数。
//
// 在一步操作中，你可以使数组中的一个元素加 1 或者减 1 。

func minMoves2(nums []int) (ans int) {

	sort.Ints(nums)
	avg := nums[len(nums)/2]

	for _, num := range nums {
		ans += absFunc(avg - num)
	}

	return ans
}

func absFunc(num int) int {
	if num >= 0 {
		return num
	} else {
		return -num
	}
}
