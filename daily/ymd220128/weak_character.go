package ymd220128

import (
	"sort"
)

// 你正在参加一个多角色游戏，每个角色都有两个主要属性：攻击 和 防御 。给你一个二维整数数组 properties ，其中 properties[i] = [attacki, defensei] 表示游戏中第 i 个角色的属性。
//
// 如果存在一个其他角色的攻击和防御等级 都严格高于 该角色的攻击和防御等级，则认为该角色为 弱角色 。
// 更正式地，如果认为角色 i 弱于 存在的另一个角色 j ，那么 attackj > attacki 且 defensej > defensei 。
//
// 返回 弱角色 的数量。

func numberOfWeakCharacters(properties [][]int) (ans int) {
	// 按照 攻击力 排序
	sort.Slice(properties, func(i, j int) bool {
		p, q := properties[i], properties[j]
		return p[0] > q[0] || p[0] == q[0] && p[1] < q[1]
	})

	// 找到最大防御力，如果防御力低于最大值，那么这个角色必然是弱角色
	maxDef := 0
	for _, p := range properties {
		if p[1] < maxDef {
			ans++
		} else {
			maxDef = p[1]
		}
	}
	return ans
}
