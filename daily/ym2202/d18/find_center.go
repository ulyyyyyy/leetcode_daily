package d18

func findCenter(edges [][]int) (node int) {
	if len(edges) < 2 {
		return
	}

	for i := 0; i < len(edges[0]); i++ {
		for j := 0; j < len(edges[1]); j++ {
			if edges[0][i] == edges[1][j] {
				return edges[0][i]
			}
		}
	}
	return
}

func findCenter2(edges [][]int) (node int) {
	if edges[0][0] == edges[1][0] || edges[0][0] == edges[1][1] {
		return edges[0][0]
	}
	return edges[0][1]
}
