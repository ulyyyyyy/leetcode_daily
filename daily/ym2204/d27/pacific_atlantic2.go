package d27

type pos struct {
	x int
	y int
}

func pacificAtlantic2(heights [][]int) (ans [][]int) {
	m, n := len(heights), len(heights[0])

	p := make(map[pos]struct{})
	a := make(map[pos]struct{})

	var dfs func(x, y int, ocean map[pos]struct{})
	dfs = func(x, y int, ocean map[pos]struct{}) {
		tPos := pos{
			x: x, y: y,
		}
		if _, ok := ocean[tPos]; ok {
			return
		}
		ocean[tPos] = struct{}{}
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

	for k := range p {
		if _, ok := a[k]; ok {
			ans = append(ans, []int{k.x, k.y})
		}
	}
	return
}
