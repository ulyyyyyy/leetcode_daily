package p3

import "sort"

func minimumLines(stockPrices [][]int) (ans int) {
	if len(stockPrices) <= 2 {
		return len(stockPrices) - 1
	}

	sort.Slice(stockPrices, func(i, j int) bool {
		return stockPrices[i][0] < stockPrices[j][0]
	})
	pre := stockPrices[1]

	slop := calcSlop(stockPrices[0], pre) // 斜率
	ans = 1

	for i := 2; i < len(stockPrices); i++ {
		last := stockPrices[i]

		nowSlop := calcSlop(pre, last)
		if nowSlop != slop {
			ans++
			slop = nowSlop
		}
		pre = last
	}

	return ans
}

func calcSlop(pos1, pos2 []int) float64 {
	x1, y1 := pos1[0], pos1[1]

	x2, y2 := pos2[0], pos2[1]

	return float64(y2-y1) / float64(x2-x1)
}
