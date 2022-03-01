package d01

// convert2 压缩矩阵，忽略无需覆盖的区域，极大减少空间与时间开销。
// 时间复杂度 O(n)
// 空间复杂度 O(n)
func convert2(s string, numRows int) (ret string) {
	if numRows == 1 || numRows > len(s) {
		return s
	}
	grid := make([][]string, numRows)
	for i := range grid {
		grid[i] = make([]string, 0)
	}

	flag := false
	next := func(preC, preR *int) {
		// 首先判断是否到顶边
		if (*preC) == 0 || (*preC) == numRows-1 {
			flag = !flag
		}
		if flag {
			*preC++
		} else {
			*preC--
			*preR++
		}
	}

	preColumn, preRow := 0, 0
	for _, c := range s {
		grid[preColumn] = append(grid[preColumn], string(c))
		next(&preColumn, &preRow)
	}

	for _, g := range grid {
		for _, v := range g {
			ret += v
		}
	}

	return ret
}
