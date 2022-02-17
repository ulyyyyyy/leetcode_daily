package wk280

import "sort"

func minimumRemoval(beans []int) int64 {
	sort.Ints(beans)
	total := 0
	for _, v := range beans {
		total += v
	}
	res := total
	n := len(beans)
	var tmp int
	for i, bean := range beans {
		tmp = total - bean*(n-i)
		if tmp < res {
			res = tmp
		}
	}
	return int64(res)
}
