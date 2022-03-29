package d29

// 一位老师正在出一场由 n道判断题构成的考试，每道题的答案为 true （用 'T' 表示）或者 false （用 'F'表示）。老师想增加学生对自己做出答案的不确定性，方法是最大化有 连续相同结果的题数。（也就是连续出现 true 或者连续出现 false）。
//
//给你一个字符串answerKey，其中answerKey[i]是第 i个问题的正确结果。除此以外，还给你一个整数 k，表示你能进行以下操作的最多次数：
//
//每次操作中，将问题的正确答案改为'T' 或者'F'（也就是将 answerKey[i] 改为'T'或者'F'）。
//请你返回在不超过 k次操作的情况下，最大连续 'T'或者 'F'的数目。

// maxConsecutiveAnswers
func maxConsecutiveAnswers(answerKey string, k int) (max int) {
	size := len(answerKey)
	if k >= size {
		return size
	}
	getMax := func(char byte) (max int) {
		left, sum := 0, 0
		for right := range answerKey {
			// 如果值不为指定的 byte，那么统计个数加一
			if answerKey[right] != char {
				sum++
			}
			// 往回找，剔除不为 byte 的元素
			for sum > k {
				if answerKey[left] != char {
					sum--
				}
				left++
			}
			// 记录最大值
			max = maxFunc(max, right-left+1)
		}
		return
	}
	return maxFunc(getMax('T'), getMax('F'))
}

func maxFunc(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
