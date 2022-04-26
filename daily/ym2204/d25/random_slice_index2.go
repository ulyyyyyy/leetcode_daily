package d25

import "math/rand"

// 抽水池采样方法。
// 这个方法不用预先处理数据，直接使用原始数据。
// 在选择数据时候再遍历
// 这种方法将时间复杂度提升到了 O(C * n)，其中 C 表示操作次数
// 更适用于预处理复杂，操作次数少的类型

// Solution2
type Solution2 []int

func Constructor2(nums []int) Solution2 {
	return nums
}

func (s Solution2) Pick2(target int) (ans int) {
	cnt := 0
	for i, num := range s {
		if num == target {
			cnt++ // 第 cnt 次遇到 target
			if rand.Intn(cnt) == 0 {
				ans = i
			}
		}
	}
	return
}
