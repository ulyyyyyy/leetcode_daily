package d07

import "sort"

// 如果字符串中不含有任何 'aaa'，'bbb' 或 'ccc' 这样的字符串作为子串，那么该字符串就是一个「快乐字符串」。
//
// 给你三个整数 a，b ，c，请你返回 任意一个 满足下列全部条件的字符串 s：
//
// s 是一个尽可能长的快乐字符串。
// s 中 最多 有a 个字母 'a'、b 个字母 'b'、c 个字母 'c' 。
// s 中只含有 'a'、'b' 、'c' 三种字母。
// 如果不存在这样的字符串 s ，请返回一个空字符串 ""。

func longestDiverseString(a, b, c int) string {
	var ans []byte
	cnt := []struct {
		c  int
		ch byte
	}{{a, 'a'}, {b, 'b'}, {c, 'c'}}
	for {
		sort.Slice(cnt, func(i, j int) bool { return cnt[i].c > cnt[j].c })
		hasNext := false
		for i, p := range cnt {
			if p.c <= 0 {
				break
			}
			m := len(ans)
			if m >= 2 && ans[m-2] == p.ch && ans[m-1] == p.ch {
				continue
			}
			hasNext = true
			ans = append(ans, p.ch)
			cnt[i].c--
			break
		}
		if !hasNext {
			return string(ans)
		}
	}
}
