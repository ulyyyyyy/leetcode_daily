package ymd220205

// 你要开发一座金矿，地质勘测学家已经探明了这座金矿中的资源分布，并用大小为 m * n 的网格 grid 进行了标注。每个单元格中的整数就表示这一单元格中的黄金数量；如果该单元格是空的，那么就是 0。
//
// 为了使收益最大化，矿工需要按以下规则来开采黄金：
//
// 每当矿工进入一个单元，就会收集该单元格中的所有黄金。
// 矿工每次可以从当前位置向上下左右四个方向走。
// 每个单元格只能被开采（进入）一次。
// 不得开采（进入）黄金数目为 0 的单元格。
// 矿工可以从网格中 任意一个 有黄金的单元格出发或者是停止。

var (
	dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
)

func getMaximumGold(grid [][]int) (ans int) {

	// 定义接口方法
	var dfs func(x, y, gold int)
	dfs = func(x, y, gold int) {
		// 更新当前值
		gold += grid[x][y]
		// 最大值更新
		if gold > ans {
			ans = gold
		}

		// 走过的路认为将金子挖完了
		tmp := grid[x][y]
		grid[x][y] = 0
		for _, dir := range dirs {
			nx, ny := x+dir.x, y+dir.y
			// 边界判断
			if (nx >= 0 && nx < len(grid)) && (ny >= 0 && ny < len(grid[0])) && grid[nx][ny] > 0 {
				dfs(nx, ny, gold)
			}
		}
		// 还原初始值
		grid[x][y] = tmp
	}

	for i := range grid {
		for j := range grid[i] {
			// 如果该点不为0，那么可以作为起点遍历
			if grid[i][j] > 0 {
				dfs(i, j, 0)
			}
		}
	}
	return
}
