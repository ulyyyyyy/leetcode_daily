package min_steps

func minSteps(s string, t string) int {
	countS := countEle(s)

	countT := countEle(t)
	tmpMap := make(map[string]int)

	for n, c := range countS {
		tmpMap[n] = maxFunc(tmpMap[n], c)
	}
	for n, c := range countT {
		tmpMap[n] = maxFunc(tmpMap[n], c)
	}
	total := 0
	for _, c := range tmpMap {
		total += c
	}
	return (total - len(s)) + (total - len(t))
}

func countEle(str string) (ret map[string]int) {
	ret = make(map[string]int)
	for _, r := range str {
		ret[string(r)]++
	}
	return
}

func maxFunc(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
