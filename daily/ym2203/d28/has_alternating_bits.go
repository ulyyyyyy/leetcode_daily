package d28

// 给定一个正整数，检查它的二进制表示是否总是 0、1 交替出现：换句话说，就是二进制表示中相邻两位的数字永不相同。

// hasAlternatingBits 利用辗转相除法逐位获取，判断与上一位是否相同
func hasAlternatingBits(n int) bool {
	if n < 2 && n >= 0 {
		return true
	}
	pre := -1
	for n >= 2 {
		t := n % 2
		if t == pre {
			return false
		}
		pre = t
		n /= 2
	}
	return pre != n
}

// hasAlternatingBits2 位运算法
func hasAlternatingBits2(n int) bool {
	// 移位运算，右移1，进行错位异或
	// 例如： 5 ^ 5>>1
	// 表现为:  101 ^ 010，如果数为交替二进制数，那么结果必全为 1
	xor := (n >> 1) ^ n
	// 此时 xor 与 xor+1 进行 与 操作
	// 如果 xor 全为1，那么 xor + 1 必然为 1 后接 n 个 0
	// 此时进行与操作，那么结果必然为 0，所以只要判断最后结果是否为 0 即可。
	rlt := (xor & (xor + 1)) == 0
	return rlt
}
