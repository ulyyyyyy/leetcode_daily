package d25

import "math/rand"

type Solution struct {
	m map[int][]int
}

func Constructor(nums []int) Solution {
	var m map[int][]int
	for idx, n := range nums {
		if _, ok := m[n]; ok {
			m[n] = append(m[n], idx)
		} else {
			m[n] = []int{idx}
		}
	}
	return Solution{
		m: m,
	}
}

func (s *Solution) Pick(target int) int {
	return getRandomEle(s.m[target])
}

func getRandomEle(slice []int) int {
	return slice[rand.Intn(len(slice))]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
