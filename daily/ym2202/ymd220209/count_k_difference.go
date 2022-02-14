package ymd220209

// 给你一个整数数组 nums 和一个整数 k ，请你返回数对 (i, j) 的数目，满足 i < j 且 |nums[i] - nums[j]| == k 。
//
// |x| 的值定义为：
//
// 如果 x >= 0 ，那么值为 x 。
// 如果 x < 0 ，那么值为 -x 。

// countKDifference 先转map，然后找到差额为k的数，累加
func countKDifference(nums []int, k int) (ans int) {
	tmp := make(map[int]int)
	for _, n := range nums {
		tmp[n]++
	}

	d := []int{1, -1}
	for num, v := range tmp {
		for _, t := range d {
			b := num + t*k
			if vv, ok := tmp[b]; ok {
				ans += v * vv
			}
		}
	}
	return ans / 2
}
