package ymd220126

// 给你一个在 X-Y 平面上的点构成的数据流。设计一个满足下述要求的算法：
//
// 添加 一个在数据流中的新点到某个数据结构中。可以添加 重复 的点，并会视作不同的点进行处理。
// 给你一个查询点，请你从数据结构中选出三个点，使这三个点和查询点一同构成一个 面积为正 的 轴对齐正方形 ，统计 满足该要求的方案数目。
// 轴对齐正方形 是一个正方形，除四条边长度相同外，还满足每条边都与 x-轴 或 y-轴 平行或垂直。
//
// 实现 DetectSquares 类：
//
// DetectSquares() 使用空数据结构初始化对象
// void add(int[] point) 向数据结构添加一个新的点 point = [x, y]
// int count(int[] point) 统计按上述方式与点 point = [x, y] 共同构造 轴对齐正方形 的方案数。

type DetectSquares struct {
	points map[int][]int
}

func Constructor() DetectSquares {
	return DetectSquares{points: make(map[int][]int, 0)}
}

func (ds *DetectSquares) Add(point []int) {
	ds.points[point[0]] = append(ds.points[point[0]], point[1])
}

func (ds *DetectSquares) Count(point []int) (count int) {
	x, y := point[0], point[1]

	allowedPoint := ds.points[x]
	if len(allowedPoint) == 0 {
		return 0
	}
	for _, p := range allowedPoint {
		if p == y {
			continue
		}

		distance := y - p
		if bp, ok := ds.points[x-distance]; ok {
			// TODO 添加逻辑
		}
	}
	return count
}

/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */
