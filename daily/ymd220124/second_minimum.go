package ymd220124

import "math"

// 城市用一个 双向连通 图表示，图中有 n 个节点，从 1 到 n 编号（包含 1 和 n）。图中的边用一个二维整数数组 edges 表示，其中每个 edges[i] = [ui, vi] 表示一条节点 ui 和节点 vi 之间的双向连通边。每组节点对由 最多一条 边连通，顶点不存在连接到自身的边。穿过任意一条边的时间是 time 分钟。
//
// 每个节点都有一个交通信号灯，每 change 分钟改变一次，从绿色变成红色，再由红色变成绿色，循环往复。所有信号灯都 同时 改变。你可以在 任何时候 进入某个节点，但是 只能 在节点 信号灯是绿色时 才能离开。如果信号灯是  绿色 ，你 不能 在节点等待，必须离开。
//
// 第二小的值 是 严格大于 最小值的所有值中最小的值。
//
// 例如，[2, 3, 4] 中第二小的值是 3 ，而 [2, 2, 4] 中第二小的值是 4 。
// 给你 n、edges、time 和 change ，返回从节点 1 到节点 n 需要的 第二短时间 。
//
// 注意：
//
// 你可以 任意次 穿过任意顶点，包括 1 和 n 。
// 你可以假设在 启程时 ，所有信号灯刚刚变成 绿色 。

func secondMinimum(n int, edges [][]int, time int, change int) int {
	// 构建图
	arcs := make([][]int, n+1)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		arcs[x] = append(arcs[x], y)
		arcs[y] = append(arcs[y], x)
	}
	// 记录1到每个节点的距离,初始值为最大值
	dist := make([][2]int, n+1)
	dist[1][1] = math.MaxInt32
	for i := 2; i <= n; i++ {
		dist[i] = [2]int{math.MaxInt32, math.MaxInt32}
	}

	type node struct {
		x int // 坐标点
		d int // 距离
	}
	// 初始点
	nowNode := []node{{1, 0}}

	// 一直遍历，找到次短路径，Dijkstra
	for dist[n][1] == math.MaxInt32 {
		p := nowNode[0]       // 得到当前点
		nowNode = nowNode[1:] // 把当前点清空

		for _, y := range arcs[p.x] {
			d := p.d + 1
			// 如果到该点的路径最短，则更新最短路径
			if d < dist[y][0] {
				// 更新最短路径
				dist[y][0] = d
				// 重置当前点
				nowNode = append(nowNode, node{y, d})
			} else if d > dist[y][0] && d < dist[y][1] { // 如果是次短路径，那么更新次短路径
				dist[y][1] = d
				nowNode = append(nowNode, node{y, d})
			}
		}
	}

	ans := 0
	for i := 0; i < dist[n][1]; i++ {
		if ans%(change*2) >= change {
			ans += change*2 - ans%(change*2)
		}
		ans += time
	}
	return ans
}
