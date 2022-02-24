package d24

func findBall(grid [][]int) (answers []int) {
	n := len(grid[0])
	ans := make([]int, n)
	for j := range ans {
		// 初始列
		col := j
		// 球下降轨迹
		for _, row := range grid {
			// 当前方向
			dir := row[col]
			// 下降
			col += dir // 移动球
			// 判断是否到达边界或卡住
			if col < 0 || col == n || row[col] != dir {
				// 如果不能到达底部，则为-1
				col = -1
				break
			}
		}
		// col >= 0 为成功到达底部
		ans[j] = col
	}
	return ans
}
