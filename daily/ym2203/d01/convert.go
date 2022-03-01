package d01

import (
	"math"
)

// convert 直接模拟，构建一个全量矩阵，将数据依次填充。
// 最后遍历矩阵，输出字符串
// 时间复杂度 O(n*numRows)
// 空间复杂度 O(n*numRows)
func convert(s string, numRows int) (ret string) {
	if numRows == 1 || numRows >= len(s) {
		return s
	}
	grid := make([][]string, numRows)
	xLen := int(math.Ceil(float64(len(s))/float64((2*numRows)-2)) * float64(numRows-1))
	for i := range grid {
		grid[i] = make([]string, xLen)
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
		grid[preColumn][preRow] = string(c)
		next(&preColumn, &preRow)
	}

	for _, g := range grid {
		for _, v := range g {
			ret += v
		}
	}

	return ret
}
