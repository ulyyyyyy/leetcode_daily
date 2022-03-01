package minimum_time

import (
	"math"
	"sort"
)

func minimumTime2(time []int, totalTrips int) (count int64) {
	sort.Ints(time)
	fast := int64(time[0])
	min := int64(math.MaxInt64)
	totalTrips64 := int64(totalTrips)
	// i 为最快一趟车跑完所需时间
	for i := fast; i >= 0; i++ {
		tmpTrips := i / fast
		// 从第二趟车开始计算可以的趟数
		for j := 1; j < len(time); j++ {
			tmpTrips += i / int64(time[j])
		}
		if tmpTrips >= totalTrips64 {
			min = minFunc(min, i)
			return min
		}
	}
	return min
}
