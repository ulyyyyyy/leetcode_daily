package d03

func addDigits(num int) int {
	// 一直遍历，直到结果小于10
	for num >= 10 {
		sum := 0
		// 逐位累加，将最终值返回给num
		for ; num > 0; num /= 10 {
			sum += num % 10
		}
		num = sum
	}
	return num
}
