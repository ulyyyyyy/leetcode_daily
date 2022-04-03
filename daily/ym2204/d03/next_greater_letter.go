package d03

import "sort"

// 给你一个排序后的字符列表 letters ，列表中只包含小写英文字母。另给出一个目标字母 target，请你寻找在这一有序列表里比目标字母大的最小字母。
//
// 在比较时，字母是依序循环出现的。举个例子：
//
// 如果目标字母 target = 'z' 并且字符列表为 letters = ['a', 'b']，则答案返回 'a'

// nextGreatestLetter 二分查找
func nextGreatestLetter(letters []byte, target byte) byte {
	size := len(letters)
	if target > letters[size-1] {
		return letters[0]
	}
	return letters[sort.Search(size-1, func(i int) bool { return letters[i] > target })]
}

// nextGreatestLetter2 线性查找
func nextGreatestLetter2(letters []byte, target byte) byte {
	for _, r := range letters {
		if r > target {
			return r
		}
	}
	return letters[0]
}
