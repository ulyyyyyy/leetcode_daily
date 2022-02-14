package ymd220203

// 给你数字 k ，请你返回和为 k 的斐波那契数字的最少数目，其中，每个斐波那契数字都可以被使用多次。
//
// 斐波那契数字定义为：
//
// F1 = 1
// F2 = 1
// Fn = Fn-1 + Fn-2 ， 其中 n > 2 。
// 数据保证对于给定的 k ，一定能找到可行解。

func findMinFibonacciNumbers(k int) (ans int) {
	// 构建斐波那契数列
	f := []int{1, 1}
	for f[len(f)-1] < k {
		f = append(f, f[len(f)-1]+f[len(f)-2])
	}

	// 倒序贪心获取
	for i := len(f) - 1; k > 0; i-- {
		if k >= f[i] {
			k -= f[i]
			ans++
		}
	}
	return
}
