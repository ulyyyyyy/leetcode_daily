package d13

func oneEditAway(first string, second string) bool {
	fLen, sLen := len(first), len(second)

	switch fLen - sLen {
	case 1:
		return isDelete(first, second)
	case 0:
		return isEdit(first, second)
	case -1:
		// second > first, mean DELETE
		return isInsert(first, second)
	default:
		return false
	}
}

// check if there is an EDIT operation between first string and second string
// len(first) == len(second)
func isEdit(first, second string) bool {
	cnt := 0
	l := len(first)

	for i := 0; i < l; i++ {
		if first[i] != second[i] {
			cnt++
			if cnt == 2 {
				return false
			}
		}
	}
	return true
}

// check if there is an INSERT operation between first and second string
// len(first) less than len(second)
// len(first) < len(second)
func isInsert(first, second string) bool {
	fLen := len(first)
	sLen := len(second)
	cnt := 0

	// Iterate over the strings and compare bit by bit
	for i, j := 0, 0; i < sLen && j < fLen; i, j = i+1, j+1 {
		if second[i] != first[j] {
			cnt++
			j--
			if cnt == 2 {
				return false
			}
		}
	}
	return true
}

// isDelete is the same as reversed var isInsert
func isDelete(first, second string) bool {
	return isInsert(second, first)
}
