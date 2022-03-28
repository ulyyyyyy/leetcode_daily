package d07

import (
	"math"
	"strconv"
)

func ConvertToBase(num int, r int) (ans string) {
	if num < r && num >= 0 {
		return strconv.Itoa(num)
	}
	flag := num < 0
	num = int(math.Abs(float64(num)))

	for num >= r {
		t := num % r
		ans = strconv.Itoa(t) + ans
		num /= r
	}
	ans = strconv.Itoa(num) + ans
	if flag {
		ans = "-" + ans
	}
	return ans
}
