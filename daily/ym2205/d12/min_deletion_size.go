package d12

// 给你由 n 个小写字母字符串组成的数组 strs，其中每个字符串长度相等。
//
// 这些字符串可以每个一行，排成一个网格。例如，strs = ["abc", "bce", "cae"] 可以排列为：
//
// abc
// bce
// cae
// 你需要找出并删除 不是按字典序升序排列的 列。在上面的例子（下标从 0 开始）中，列 0（'a', 'b', 'c'）和列 2（'c', 'e', 'e'）都是按升序排列的，而列 1（'b', 'c', 'a'）不是，所以要删除列 1 。
//
// 返回你需要删除的列数。

// minDeletionSize 找到每列字符串顺序，使用常数空间 pre 记录上一个字符，判断是否为升序顺序。
// 如果非升序顺序，直接 ans++，跳过该列
func minDeletionSize(strs []string) (ans int) {
	// 如果每个字符串元素个数少于 1，直接返回 0
	l := len(strs[0])
	if l == 0 {
		return
	}
	for i := 0; i < l; i++ {
		// 初始值为最小值
		pre := byte('a') - 1
		for j := 0; j < len(strs); j++ {
			// 判断是否升序
			if strs[j][i] >= pre {
				pre = strs[j][i]
			} else {
				ans++
				break
			}
		}
	}
	return
}
