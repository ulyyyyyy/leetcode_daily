package d18

// 给你一个整数 n ，按字典序返回范围 [1, n] 内所有整数。
//
// 你必须设计一个时间复杂度为 O(n) 且使用 O(1) 额外空间的算法。

// lexicalOrder2 尝试写时间复杂度O(n)空间复杂度O(1)的算法
// 递归改为迭代
func lexicalOrder2(n int) (slice []int) {
	slice = make([]int, n)
	num := 1
	for i := range slice {
		slice[i] = num
		if num*10 <= n {
			num *= 10
		} else {
			for num%10 == 9 || num+1 > n {
				num /= 10
			}
			num++
		}
	}
	return slice
}
