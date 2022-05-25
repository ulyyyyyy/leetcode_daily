package p4

import (
	"math"
)

func totalStrength(strength []int) (ans int) {
	n := len(strength)
	tmpAns := int64(0)

	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			min, total := minAndTotal(strength[i:j])

			tmpAns += min * total
		}
	}

	ans = int(tmpAns % int64(math.Pow(10.0, 9.0)+7))
	return
}

func minAndTotal(nums []int) (min, total int64) {
	min = math.MaxInt64

	minFunc := func(a, b int64) int64 {
		if a >= b {
			return b
		}
		return a
	}

	for _, num := range nums {
		tn := int64(num)
		min = minFunc(tn, min)
		total += tn
	}
	return
}
