package minimum_time

import (
	"math"
	"sort"
)

func minimumTime(time []int, totalTrips int) (count int64) {
	sort.Ints(time)
	fast := int64(time[0])
	min := fast
	totalTrips64 := int64(totalTrips)

	// pre 为最快一趟车跑完所需时间
	pre := fast
	last := totalTrips64 * fast

	calcTrips := func(mid int64) (tmpTrips int64) {
		for _, t := range time {
			if int64(t) <= mid {
				tmpTrips += mid / int64(t)
			}
		}
		return
	}

	for pre < last {
		mid := int64(math.Ceil(float64(pre+last) / 2))
		tmpTrips := calcTrips(mid)

		if tmpTrips < totalTrips64 {
			pre = mid
			continue
		}

		if last-pre > 1 {
			last = mid
		} else if calcTrips(pre) == totalTrips64 {
			return pre
		} else {
			return last
		}
	}
	return min
}
func minFunc(a, b int64) int64 {

	if a >= b {
		return b
	}
	return a
}
