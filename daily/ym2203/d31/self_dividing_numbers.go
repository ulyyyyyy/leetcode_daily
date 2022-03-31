package d31

// 自除数 是指可以被它包含的每一位数整除的数。
//
// 例如，128 是一个 自除数 ，因为 128 % 1 == 0，128 % 2 == 0，128 % 8 == 0。
// 自除数 不允许包含 0 。
//
// 给定两个整数 left 和 right ，返回一个列表，列表的元素是范围 [left, right] 内所有的 自除数 。

func selfDividingNumbers(left int, right int) (slice []int) {
outer:
	for i := left; i <= right; i++ {
		tmp := i
		for tmp > 0 {
			mod := tmp % 10
			if mod == 0 || i%mod != 0 {
				continue outer
			}
			tmp /= 10
		}
		slice = append(slice, i)
	}
	return
}
