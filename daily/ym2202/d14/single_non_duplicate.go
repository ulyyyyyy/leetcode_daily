package d14

// 给你一个仅由整数组成的有序数组，其中每个元素都会出现两次，唯有一个数只会出现一次。
//
// 请你找出并返回只出现一次的那个数。
//
// 你设计的解决方案必须满足 O(log n) 时间复杂度和 O(1) 空间复杂度。

func singleNonDuplicate(nums []int) (ans int) {

	l := len(nums)
	mid := func() (rlt bool, c int) {
		for i := 0; i < l/2; i += 2 {
			c++
			if nums[i] != nums[i+1] {
				return true, nums[i]
			}
		}
		nums = nums[c*2:]
		return false, ans
	}

	for len(nums) != 1 {
		rlt, a := mid()
		if rlt {
			return a
		}
	}
	return nums[0]
}
