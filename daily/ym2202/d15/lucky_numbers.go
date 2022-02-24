package d15

func luckyNumbers(matrix [][]int) (slice []int) {
	m := len(matrix)

	findMinIndex := func(raw []int) (index int) {
		min := raw[0]
		for i, v := range raw {
			if v < min {
				min = v
				index = i
			}
		}
		return
	}

Outer:
	for x, raw := range matrix {
		y := findMinIndex(raw)

		for i := 0; i < m; i++ {
			if matrix[i][y] > matrix[x][y] {
				continue Outer
			}
		}
		slice = append(slice, matrix[x][y])
	}
	return
}
