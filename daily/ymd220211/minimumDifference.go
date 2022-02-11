package ymd220211

import (
	"math"
	"sort"
)

// 给你一个 下标从 0 开始 的整数数组 nums ，其中 nums[i] 表示第 i 名学生的分数。另给你一个整数 k 。
//
// 从数组中选出任意 k 名学生的分数，使这 k 个分数间 最高分 和 最低分 的 差值 达到 最小化 。
//
// 返回可能的 最小差值 。

// minimumDifference 排序后，利用固定滑动窗口来解决
func minimumDifference(nums []int, k int) (ans int) {
	if k == 1 {
		return 0
	}
	ans = math.MaxInt
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		j := i + k - 1
		if j < len(nums) && ans > nums[j]-nums[i] {
			ans = nums[j] - nums[i]
		}
	}
	return
}
