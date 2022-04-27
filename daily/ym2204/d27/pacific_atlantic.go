package d27

var directs = [][]int{
	{-1, 0}, // 上
	{1, 0},  // 下
	{0, -1}, // 左
	{0, 1},  // 右
}

func pacificAtlantic(heights [][]int) (ans [][]int) {
	m, n := len(heights), len(heights[0])

	p := make([][]bool, m)
	a := make([][]bool, m)

	for i := 0; i < m; i++ {
		p[i] = make([]bool, n)
		a[i] = make([]bool, n)
	}

	var dfs func(x, y int, ocean [][]bool)
	dfs = func(x, y int, ocean [][]bool) {
		if ocean[x][y] {
			return
		}
		ocean[x][y] = true
		// 检查当前节点是否为边界值
		for _, d := range directs {
			tx, ty := x+d[0], y+d[1]
			isValid := tx >= 0 && tx < m && ty >= 0 && ty < n
			// 如果从边界能够蔓延过来，那么就增加该节点
			if isValid && heights[x][y] <= heights[tx][ty] {
				dfs(tx, ty, ocean)
			}
		}
	}

	for x := 0; x < m; x++ {
		// 太平洋
		dfs(x, 0, p)
		// 大西洋
		dfs(x, n-1, a)
	}

	for y := 0; y < n; y++ {
		dfs(0, y, p)
		dfs(m-1, y, a)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if p[i][j] && a[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return
}
