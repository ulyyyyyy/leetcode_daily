package d18

// 给你一个整数 n ，按字典序返回范围 [1, n] 内所有整数。
//
// 你必须设计一个时间复杂度为 O(n) 且使用 O(1) 额外空间的算法。
//
// 字典序： 1, 10, 100, 110

// lexicalOrder 直接递归
func lexicalOrder(n int) (slice []int) {
	var dfs func(num int)
	dfs = func(num int) {
		num *= 10
		for i := 0; i < 10; i++ {
			if num+i <= n {
				slice = append(slice, num+i)
			} else {
				return
			}
			dfs(num + i)
		}
	}

	for i := 1; i < 10; i++ {
		if i <= n {
			slice = append(slice, i)
			dfs(i)
		}
	}
	return slice
}
