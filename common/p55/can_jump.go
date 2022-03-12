package p55

// 给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
//
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
// 判断你是否能够到达最后一个下标。

// canJump 跳跃游戏，可以看作简单的贪心算法。
// 维护一个最大值，每次跳跃都更新最大值。
// 如果中间卡住了则直接返回 false
func canJump(nums []int) bool {
	max := 0
	for i, num := range nums {
		// 判断最远距离能不能到当前节点，如果不能就直接退出
		if max < i {
			return false
		}
		// 判断是否需要更新最大节点，同时判断最大距离是否超过了长度
		// 如果超过了就直接返回 true
		if i+num >= max {
			max = i + num
			if max >= len(nums)-1 {
				return true
			}
		}
	}
	// 说明更新的最大值始终没有到达最远节点，那么直接返回 false
	return false
}
