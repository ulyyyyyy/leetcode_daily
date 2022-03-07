package d07

import (
	"math"
	"strconv"
)

func convertToBase7(num int) (ans string) {
	if num < 7 && num >= 0 {
		return strconv.Itoa(num)
	}
	flag := num < 0
	num = int(math.Abs(float64(num)))

	for num >= 7 {
		t := num % 7
		ans = strconv.Itoa(t) + ans
		num /= 7
	}
	ans = strconv.Itoa(num) + ans
	if flag {
		ans = "-" + ans
	}
	return ans
}
