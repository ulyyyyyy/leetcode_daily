package d11

// countHighestScoreNodesOffice 思路，以当前树为基础，依次拿掉每个节点，计算其最大值情况。
// 最后返回最大值的次数。
func countHighestScoreNodesOffice(parents []int) (ans int) {
	n := len(parents)
	children := make([][]int, n)
	// 构建每个节点下的子节点情况
	for node := 1; node < n; node++ {
		p := parents[node]
		children[p] = append(children[p], node)
	}
	// 维护最大值
	maxScore := 0
	// 深度遍历方法
	var dfs func(node int) int
	dfs = func(node int) int {
		// 维护每一个节点的得分与节点总数
		score, size := 1, n-1
		// 深度优先遍历每个节点
		for _, child := range children[node] {
			sz := dfs(child)
			score *= sz
			size -= sz
		}
		// 如果
		if node > 0 {
			score *= size
		}
		if score == maxScore {
			ans++
		} else if score > maxScore {
			ans = 1
			maxScore = score
		}
		return n - size
	}
	dfs(0)
	return
}
