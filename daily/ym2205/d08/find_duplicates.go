package d08

// 给你一个长度为 n 的整数数组 nums ，其中 nums 的所有整数都在范围 [1, n] 内，且每个整数出现 一次 或 两次 。
//
// 请你找出所有出现 两次 的整数，并以数组形式返回。
//
// 你必须设计并实现一个时间复杂度为 O(n) 且仅使用常量额外空间的算法解决此问题。

// findDuplicates
func findDuplicates(nums []int) (ans []int) {
	l := len(nums)
	tmp := make([]int, l)

	// 将原列表数据与其下表做对应
	for _, n := range nums {
		// 说明不存在该元素
		if tmp[n-1] == 0 {
			tmp[n-1] = n
		} else {
			tmp[n-1] = -1
			ans = append(ans, n)
		}
	}

	return ans
}
