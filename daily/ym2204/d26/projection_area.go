package d26

func projectionArea(grid [][]int) int {
	// 找到 xy 面的投影，即从z轴观察
	// 投影面积等于 x,y 面的个数，即 grid 中单元格的个数
	Sxy := 0

	Sxz := 0
	Syz := 0

	l := len(grid)
	maxX := make([]int, l)
	maxY := make([]int, l)
	for idx := range grid {
		for i, s := range grid[idx] {
			if s != 0 {
				Sxy += 1
			}
			maxX[idx] = maxFunc(maxX[idx], s)
			maxY[i] = maxFunc(maxY[i], s)
		}
	}

	for _, n := range maxX {
		Syz += n
	}

	for _, n := range maxY {
		Sxz += n
	}
	return Sxz + Sxy + Syz
}

func maxFunc(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
