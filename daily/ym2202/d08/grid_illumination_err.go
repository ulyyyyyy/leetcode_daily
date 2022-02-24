package d08

// 在大小为 n x n 的网格 grid 上，每个单元格都有一盏灯，最初灯都处于 关闭 状态。
//
// 给你一个由灯的位置组成的二维数组 lamps ，其中 lamps[i] = [rowi, coli] 表示 打开 位于 grid[rowi][coli] 的灯。即便同一盏灯可能在 lamps 中多次列出，不会影响这盏灯处于 打开 状态。
//
// 当一盏灯处于打开状态，它将会照亮 自身所在单元格 以及同一 行 、同一 列 和两条 对角线 上的 所有其他单元格 。
//
// 另给你一个二维数组 queries ，其中 queries[j] = [rowj, colj] 。对于第 j 个查询，如果单元格 [rowj, colj] 是被照亮的，则查询结果为 1 ，否则为 0 。在第 j 次查询之后 [按照查询的顺序] ，关闭 位于单元格 grid[rowj][colj] 上及相邻 8 个方向上（与单元格 grid[rowi][coli] 共享角或边）的任何灯。
//
// 返回一个整数数组 ans 作为答案， ans[j] 应等于第 j 次查询 queries[j] 的结果，1 表示照亮，0 表示未照亮。

// 第一次，这里我用了硬解去做，结果在过用例的时候 OOM 了
// 所以可以反着来，遍历的时候去找横、竖、对角线是否有灯，如果没有则为0
var directs = [][]int{
	{1, 1},   // 右下
	{1, 0},   // 右
	{1, -1},  // 右上
	{0, 1},   // 下
	{0, 0},   // 自身
	{0, -1},  // 上
	{-1, 1},  // 左下
	{-1, 0},  // 左
	{-1, -1}, // 左上
}

type grid [][]int

func newGrid(n int) grid {
	g := make([][]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, n)
	}
	return g
}

func gridIllumination(n int, lamps [][]int, queries [][]int) (ans []int) {
	for _, query := range queries {
		// 构造grid
		g := newGrid(n)
		for _, lamp := range lamps {
			x, y := lamp[0], lamp[1]
			g[x][y] = 1

			for t := 0; t < n; t++ {
				g[t][y] = 1
				g[x][t] = 1

				if x-t >= 0 && x-t < n && y-t >= 0 && y-t < n {
					g[x-t][y-t] = 1
				}
				if x-t >= 0 && x-t < n && y+t >= 0 && y+t < n {
					g[x-t][y+t] = 1
				}
				if x+t >= 0 && x+t < n && y-t >= 0 && y-t < n {
					g[x+t][y-t] = 1
				}
				if x+t >= 0 && x+t < n && y+t >= 0 && y+t < n {
					g[x+t][y+t] = 1
				}
			}
		}

		x, y := query[0], query[1]

		ans = append(ans, g[x][y])

		for _, direct := range directs {
			nx, ny := x+direct[0], y+direct[1]
			for i := 0; i < len(lamps); i++ {
				if lamps[i][0] == nx && lamps[i][1] == ny {
					lamps = append(lamps[:i], lamps[i+1:]...)
					i--
				}
			}
		}
	}
	return
}
