package p11

// maxArea 首先需要明确一点，挡板不占空间...
// 然后根据公式 S = min(h[i], h[j]) * (j-i)，需要尽可能构造 j-i & min 最大值
// 所以先从两端逐步缩小（保证 j-i 最大），再找到 min 的最值
func maxArea(height []int) int {
	max := 0

	size := len(height)
	for i, j := 0, size-1; i < j; {
		max = maxFunc(max, minFunc(height[i], height[j])*(j-i))
		if height[i] >= height[j] {
			j--
		} else {
			i++
		}
	}
	return max
}

func minFunc(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func maxFunc(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
