package d20

import "sort"

// 给你一个区间数组 intervals ，其中 intervals[i] = [starti, endi] ，且每个 starti 都 不同 。
//
// 区间 i 的 右侧区间 可以记作区间 j ，并满足 startj >= endi ，且 startj 最小化 。
//
// 返回一个由每个区间 i 的 右侧区间 的最小起始位置组成的数组。如果某个区间 i 不存在对应的 右侧区间 ，则下标 i 处的值设为 -1 。

// findRightInterval
func findRightInterval(intervals [][]int) (ans []int) {
	ans = make([]int, 0, len(intervals))

	sortedIntervals := make([]int, 0, len(intervals))
	// map, start -> index
	indexMap := make(map[int]int, len(intervals))

	for idx, interval := range intervals {
		sortedIntervals = append(sortedIntervals, interval[0])
		indexMap[interval[0]] = idx
	}
	sort.Ints(sortedIntervals)

	for _, interval := range intervals {
		target := interval[1]
		idx := binarySearch(sortedIntervals, target)

		// 不存在
		if idx == -1 {
			ans = append(ans, -1)
		} else {
			ans = append(ans, indexMap[sortedIntervals[idx]])
		}
	}
	return ans
}

func binarySearch(nums []int, target int) (idx int) {
	idx = -1
	left, right := -1, len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			idx = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return idx
}
