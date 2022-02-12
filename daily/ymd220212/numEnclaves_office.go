package ymd220212

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func numEnclavesOffice(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	var dfs func(int, int)
	dfs = func(r, c int) {
		if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] == 0 || vis[r][c] {
			return
		}
		vis[r][c] = true
		for _, d := range dirs {
			dfs(r+d.x, c+d.y)
		}
	}
	// 遍历上、下两个边界的陆地
	for i := range grid {
		dfs(i, 0)
		dfs(i, n-1)
	}
	// 遍历左、右两个边界的陆地
	for j := 1; j < n-1; j++ {
		dfs(0, j)
		dfs(m-1, j)
	}
	// 遍历仅存的陆地
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 1 && !vis[i][j] {
				ans++
			}
		}
	}
	return
}

// 作者：LeetCode-Solution
// 链接：https://leetcode-cn.com/problems/number-of-enclaves/solution/fei-di-de-shu-liang-by-leetcode-solution-nzs3/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
