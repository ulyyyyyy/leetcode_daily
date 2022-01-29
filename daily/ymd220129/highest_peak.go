package ymd220129

// 给你一个大小为 m x n 的整数矩阵 isWater ，它代表了一个由 陆地 和 水域 单元格组成的地图。
//
// 如果 isWater[i][j] == 0 ，格子 (i, j) 是一个 陆地 格子。
// 如果 isWater[i][j] == 1 ，格子 (i, j) 是一个 水域 格子。
// 你需要按照如下规则给每个单元格安排高度：
//
// 每个格子的高度都必须是非负的。
// 如果一个格子是是 水域 ，那么它的高度必须为 0 。
// 任意相邻的格子高度差 至多 为 1 。当两个格子在正东、南、西、北方向上相互紧挨着，就称它们为相邻的格子。（也就是说它们有一条公共边）
// 找到一种安排高度的方案，使得矩阵中的最高高度值 最大 。
//
// 请你返回一个大小为 m,x,n 的整数矩阵 height ，其中 height[i][j] 是格子 (i, j) 的高度。如果有多种解法，请返回 任意一个 。

var (
	offsetX = []int{1, 0, -1, 0}
	offsetY = []int{0, 1, 0, -1}
)

func highestPeak(isWater [][]int) [][]int {
	var q queue

	// 遍历初始数据，初始化数据
	for i := range isWater {
		for j := range isWater[i] {
			// 判断是否为水面，如果是水面，那么置零，同时将坐标记录队列
			// 如果不是水面，那么将高度置 -1，防止与水面高度冲突
			if isWater[i][j] == 1 {
				q.push(newPoint(i, j))
				isWater[i][j] = 0
			} else {
				isWater[i][j] = -1
			}
		}
	}
	// 记录已入库数据
	existNum := len(q)
	// 循环，直到查询完所有的地块
	for q.len() != 0 {
		for i := 0; i < q.len(); i++ {
			n := q.pop() // 取出这次要遍历的数据

			x, y := n[0], n[1] // 保存数据

			// 遍历下、右、上、左，至于为什么是这个顺序，全凭心情。
			for j := 0; j < 4; j++ {
				a, b := x+offsetX[j], y+offsetY[j]
				if a < 0 || b < 0 || a >= len(isWater) || b >= len(isWater[0]) {
					continue
				}
				// 如果这个地块还没有遍历到，那么高度相较于前一块 +1，同时放入下次遍历顺序
				if isWater[a][b] == -1 {
					isWater[a][b] = isWater[x][y] + 1
					q.push(newPoint(a, b))
					existNum++
				}
			}
		}
	}
	return isWater
}

// 队列相关数据
// queue 队列保存下次遍历的坐标
type queue []point

type point [2]int

func newPoint(x, y int) point {
	return [2]int{x, y}
}

func (q *queue) push(n point) {
	*q = append(*q, n)
}

func (q *queue) pop() point {
	ret := (*q)[0]
	*q = (*q)[1:]
	return ret
}
func (q queue) len() int {
	return len(q)
}
