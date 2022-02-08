package ymd220206

// 给你一个整数数组 nums 。数组中唯一元素是那些只出现 恰好一次 的元素。
//
// 请你返回 nums 中唯一元素的 和 。

func sumOfUnique(nums []int) (count int) {
	tmap := make(map[int]int, 0)
	for _, n := range nums {
		if c, ok := tmap[n]; !ok {
			count += n
			tmap[n] = 1
		} else if c == 1 {
			count -= n
			tmap[n]++
		}
	}
	return
}
