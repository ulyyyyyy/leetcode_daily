package p1

import "math"

func percentageLetter(s string, letter byte) int {
	letterNum := 0
	totalNum := 0
	for _, r := range s {
		if letter == byte(r) {
			letterNum++
		}
		totalNum++
	}

	return int(math.Floor(float64(letterNum) / float64(totalNum) * 100))
}
