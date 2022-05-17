package d17

import "sort"

// isAlienSorted
func isAlienSorted(words []string, order string) bool {
	orderMap := make(map[byte]int, len(order))
	// 遍历 orderMap, 保存各字符所在下标
	for idx, r := range order {
		orderMap[byte(r)] = idx
	}

	var isSortedByOrderDict func(pre, last string) bool
	// isSorted 检查两个字符串是否按照字典序排序
	isSortedByOrderDict = func(pre, last string) bool {
		if pre == "" {
			return true
		}
		if last == "" {
			return false
		}
		// 逐位比较两个字符串，判断时候符合字段序
		// 如果 pre[idx] > last[idx], 说明不满足字典序
		if orderMap[pre[0]] > orderMap[last[0]] {
			return false
		} else if orderMap[pre[0]] == orderMap[last[0]] {
			return isSortedByOrderDict(pre[1:], last[1:])
		} else {
			return true
		}
	}

	pre := words[0]
	for i := 1; i < len(words); i++ {
		last := words[i]
		// 如果没有按照字段序排序，则直接返回 false
		if !isSortedByOrderDict(pre, last) {
			return false
		}
		// 否则，维护 pre,
		pre = last
	}
	return true
}

// isAlienSorted2
func isAlienSorted2(words []string, order string) bool {
	orderMap := make(map[byte]int, len(order))
	// 遍历 orderMap, 保存各字符所在下标
	for idx, r := range order {
		orderMap[byte(r)] = idx
	}

	var isSortedByOrderDict func(pre, last string) bool
	// isSorted 检查两个字符串是否按照字典序排序
	isSortedByOrderDict = func(pre, last string) bool {
		if pre == last || pre == "" {
			return true
		}
		if last == "" {
			return false
		}
		// 逐位比较两个字符串，判断时候符合字段序
		// 如果 pre[idx] > last[idx], 说明不满足字典序
		if orderMap[pre[0]] > orderMap[last[0]] {
			return false
		} else if orderMap[pre[0]] == orderMap[last[0]] {
			return isSortedByOrderDict(pre[1:], last[1:])
		} else {
			return true
		}
	}

	return sort.SliceIsSorted(words, func(i, j int) bool {
		// 判断两个
		return isSortedByOrderDict(words[j], words[i])
	})
}
