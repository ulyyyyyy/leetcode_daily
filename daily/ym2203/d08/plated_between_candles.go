package d08

// platesBetweenCandles
func platesBetweenCandles(s string, queries [][]int) (answers []int) {
	answers = make([]int, 0, len(queries))

	// 构造蜡烛数量的下标
	platIndexSli := make([]int, 0)
	for i, r := range s {
		if r == '|' {
			platIndexSli = append(platIndexSli, i)
		}
	}

	// 找到左、右边界与中间盘子的数据，得到最后结论。
	for _, query := range queries {
		pre, last := query[0], query[1]
		//
		extraPlatesTotal := 0
		for m := range platIndexSli {
			if pre > platIndexSli[m] && pre <= platIndexSli[m+1] {
				extraPlatesTotal = 0
				pre = platIndexSli[m+1]
			}
			if last < platIndexSli[m] && last >= platIndexSli[m-1] {
				last = platIndexSli[m-1]
			}
			if platIndexSli[m] > pre && platIndexSli[m] < last {
				extraPlatesTotal++
			}
		}
		if pre >= last {
			answers = append(answers, 0)
		} else {
			answers = append(answers, last-pre-extraPlatesTotal-1)
		}
	}
	// 问题在于，如何确定左、右边界是否与extra重复
	return
}

func minFunc(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func maxFunc(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func findPlates(input string) (num int) {
	pre, last := -1, -1
	for i, v := range input {
		if pre == -1 && v == '|' {
			pre = i
		} else if pre != -1 && v == '*' {
			num++
		} else if pre != -1 && v == '|' {
			last = i
		}
	}
	if last == -1 {
		return 0
	}
	return num
}
