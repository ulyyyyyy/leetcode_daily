package d09

// bestRotation 直接模拟，超时
func bestRotation(nums []int) int {
	type kRotation struct {
		k   int
		max int
	}

	minFunc := func(m, cmpObj *kRotation) {
		// 如果发现 对比对象的最大值大于原值 且 下标值小于原值，那么更新当前状态
		if (m.max < cmpObj.max) || (m.max == cmpObj.max && cmpObj.k < m.k) {
			m.max = cmpObj.max
			m.k = cmpObj.k
		}
	}

	min := kRotation{
		k:   len(nums),
		max: 0,
	}

	// 首先按照可能的 k 值进行模拟遍历
	for k := 0; k < len(nums); k++ {
		tmpNum := append(nums[k:], nums[:k]...)

		total := 0
		for i, v := range tmpNum {
			if v <= i {
				total++
			}
		}
		minFunc(&min, &kRotation{
			k:   k,
			max: total,
		})
	}

	return min.k
}
