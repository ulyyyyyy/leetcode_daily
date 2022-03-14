package p703

import (
	"container/heap"
	"sort"
)

// 校验接口实现
var _ heap.Interface = (*KthLargest)(nil)

type KthLargest struct {
	cap int
	sort.IntSlice
}

func (k *KthLargest) Push(v interface{}) {
	k.IntSlice = append(k.IntSlice, v.(int))
}

func (k *KthLargest) Pop() interface{} {
	s := k.IntSlice
	v := s[len(s)-1]
	k.IntSlice = s[:len(s)-1]
	return v
}

func Constructor(k int, nums []int) (kth KthLargest) {
	kth = KthLargest{cap: k}
	for _, val := range nums {
		kth.Add(val)
	}
	return kth
}

func (k *KthLargest) Add(val int) int {
	heap.Push(k, val)
	if k.Len() > k.cap {
		heap.Pop(k)
	}
	return k.IntSlice[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
