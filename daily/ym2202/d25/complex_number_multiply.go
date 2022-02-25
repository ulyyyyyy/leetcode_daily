package d25

import (
	"fmt"
	"strconv"
	"strings"
)

func complexNumberMultiply(num1, num2 string) string {
	// 将字符串转为实虚数
	convert := func(num string) (real, imag int) {
		i := strings.IndexByte(num, '+')
		real, _ = strconv.Atoi(num[:i])
		imag, _ = strconv.Atoi(num[i+1 : len(num)-1])
		return
	}

	real1, imag1 := convert(num1)
	real2, imag2 := convert(num2)
	return fmt.Sprintf("%d+%di", real1*real2-imag1*imag2, real1*imag2+imag1*real2)
}

func conv(num string) (real, imag int) {
	i := strings.IndexByte(num, '+')
	real, _ = strconv.Atoi(num[:i])
	imag, _ = strconv.Atoi(num[i+1 : len(num)-1])
	return
}
func conv2(num string) (real, imag int) {
	split := strings.Split(num, "+")
	real, _ = strconv.Atoi(split[0])
	imag, _ = strconv.Atoi(split[1])
	return
}
