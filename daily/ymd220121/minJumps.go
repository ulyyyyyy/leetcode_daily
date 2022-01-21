package ymd220121

import (
	"fmt"
	"sort"
)

// NO. 1345 跳跃游戏IV
// 给你一个整数数组 arr ，你一开始在数组的第一个元素处（下标为 0）。
//
// 每一步，你可以从下标i 跳到下标：
//
// i + 1: 满足：i + 1 < arr.length
// i - 1: 满足：i - 1 >= 0
// j: 满足：arr[i] == arr[j] 且 i != j
// 请你返回到达数组最后一个元素的下标处所需的 最少操作次数 。
//
// 注意：任何时候你都不能跳到数组外面。

var (
	steps   = make([]int, 0)
	arrMap  = make(map[int][]int, 0)
	arrSize = 0
	pArr    = make([]int, 0)
	min     = 0
)

func minJumps(arr []int) int {
	pArr = arr
	generaMap()
	t := newTree(0)
	fmt.Println(t.val)
	bfs(t, len(arr)-1)
	return min
}

func bfs(t *tree, target int) {
	min++
	if t == nil {
		return
	}
	t.setChildList()
	for _, childTree := range t.childList {
		if childTree.val == target {
			return
		}
	}
	for _, childTree := range t.childList {
		bfs(childTree, target)
	}
	return
}

func generaMap() {
	for index, num := range pArr {
		arrSize++
		arrMap[num] = append(arrMap[num], index)
	}
}

func getNextIndex(index int) (allowedSteps []int) {
	steps = append(steps, index)
	allowedSteps = make([]int, 0)
	// 下一跳
	if index+1 < arrSize && !exits(index+1) {
		allowedSteps = append(allowedSteps, index+1)
	} else if index+1 == arrSize {
		return
	}
	// 上一跳
	if index-1 >= 0 && !exits(index-1) {
		allowedSteps = append(allowedSteps, index-1)
	} else if index-1 == arrSize {
		return
	}

	// 同值跳跃
	for _, sIndex := range arrMap[pArr[index]] {
		if sIndex != index && !exits(sIndex) {
			allowedSteps = append(allowedSteps, sIndex)
		}
	}
	return allowedSteps
}

// exits 判断当前下标是否已经遍历
func exits(index int) (rlt bool) {
	return sort.SearchInts(steps, index) < len(steps)
}

type tree struct {
	childList []*tree
	val       int
}

func newTree(val int) *tree {
	return &tree{val: val, childList: make([]*tree, 0)}
}

func (t *tree) setChildList() {
	for _, v := range getNextIndex(t.val) {
		t.childList = append(t.childList, newTree(v))
	}
}
