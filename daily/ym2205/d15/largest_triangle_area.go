package d15

import "math"

// 给定包含多个点的集合，从其中取三个点组成三角形，返回能组成的最大三角形的面积。

func triangleArea(x1, y1, x2, y2, x3, y3 int) float64 {
	return math.Abs(float64(x1*y2+x2*y3+x3*y1-x1*y3-x2*y1-x3*y2)) / 2
}

func largestTriangleArea(points [][]int) (ans float64) {
	for i, p := range points {
		for j, q := range points[:i] {
			for _, r := range points[:j] {
				ans = math.Max(ans, triangleArea(p[0], p[1], q[0], q[1], r[0], r[1]))
			}
		}
	}
	return
}
