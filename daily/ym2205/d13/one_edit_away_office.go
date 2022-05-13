package d13

func oneEditAway2(first, second string) bool {
	m, n := len(first), len(second)
	if m < n {
		return oneEditAway2(second, first)
	}
	if m-n > 1 {
		return false
	}
	for i, ch := range second {
		if first[i] != byte(ch) {
			if m == n {
				return first[i+1:] == second[i+1:]
			}
			return first[i+1:] == second[i:]
		}
	}
	return true
}

func compareStr(sli1, sli2 string) bool {
	return sli1 == sli2
}

func compareStrByBit(sli1, sli2 string) bool {
	l := len(sli1)
	for i := 0; i < l; i++ {
		if sli1[i] != sli2[i] {
			return false
		}
	}
	return true
}
