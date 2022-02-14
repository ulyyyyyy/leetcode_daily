package ymd220130

import "strings"

// 句子 是一串由空格分隔的单词。每个 单词 仅由小写字母组成。
//
// 如果某个单词在其中一个句子中恰好出现一次，在另一个句子中却 没有出现 ，那么这个单词就是 不常见的 。
//
// 给你两个 句子 s1 和 s2 ，返回所有 不常用单词 的列表。返回列表中单词可以按 任意顺序 组织。

func uncommonFromSentences(s1 string, s2 string) []string {
	map1 := make(map[string]int, 0)
	retSlice := make([]string, 0)
	tmpstrings := strings.Join([]string{s1, s2}, " ")
	// 遍历保存数据
	for _, word := range strings.Split(tmpstrings, " ") {
		if _, ok := map1[word]; ok {
			map1[word]++
		} else {
			map1[word] = 1
		}
	}

	for k, v := range map1 {
		if v == 1 {
			retSlice = append(retSlice, k)
		}
	}
	return retSlice
}
