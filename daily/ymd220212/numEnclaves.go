package ymd220212

// 给你一个大小为 m x n 的二进制矩阵 grid ，其中 0 表示一个海洋单元格、1 表示一个陆地单元格。
//
// 一次 移动 是指从一个陆地单元格走到另一个相邻（上、下、左、右）的陆地单元格或跨过 grid 的边界。
//
// 返回网格中 无法 在任意次数的移动中离开网格边界的陆地单元格的数量。

// 深度优点算法解决
var (
	directs = [][]int{
		{-1, 0}, // 上
		{1, 0},  // 下
		{0, -1}, // 左
		{0, 1},  // 右
	}
)

func numEnclaves(grid [][]int) (count int) {

	q := make([][]int, 0)
	var dfs func(x, y int)
	dfs = func(x, y int) {
		q = append(q, []int{x, y})
		grid[x][y] = 0
		for _, d := range directs {
			nx, ny := x+d[0], y+d[1]

			if nx >= 0 && nx < len(grid) && ny >= 0 && ny < len(grid[0]) && grid[nx][ny] == 1 {
				dfs(nx, ny)
			}
		}
	}

	// 检查是否存在与边界
	onEdge := func(x, y int) bool {
		return x >= len(grid)-1 || x < 0 || y < 0 || y >= len(grid[0])-1
	}

	for x, lineX := range grid {
		for y, v := range lineX {
			if v == 1 {
				dfs(x, y)
				tmpCount := count
				for _, p := range q {
					// 如果不在边界
					if !onEdge(p[0], p[1]) {
						count++
					} else {
						count = tmpCount
						break
					}
				}
				q = make([][]int, 0)
			}
		}
	}
	return
}
