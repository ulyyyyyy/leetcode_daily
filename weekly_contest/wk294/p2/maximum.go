package p2

import "sort"

func maximumBags(capacity []int, rocks []int, additionalRocks int) (ans int) {
	n := len(capacity)

	// 剩余数组
	needSli := make([]int, n)
	needNum := 0

	for i := 0; i < n; i++ {
		nd := capacity[i] - rocks[i]
		needSli[i] = nd
		needNum += nd
	}

	if needNum <= additionalRocks {
		return n
	}

	sort.Ints(needSli)
	for _, num := range needSli {
		if additionalRocks >= num {
			ans++
			additionalRocks -= num
		} else {
			return ans
		}
	}
	return ans
}
