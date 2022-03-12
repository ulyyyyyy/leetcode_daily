package p17

// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
var numberMap = map[string][]string{
	"":  {},
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

var retSli []string

func letterCombinations(digits string) []string {
	retSli = make([]string, 0)
	if len(digits) == 0 {
		return retSli
	}
	backTrack(digits, 0, "")
	return retSli
}

func backTrack(digits string, index int, combinations string) {
	if index != len(digits) {
		digit := string(digits[index])
		letters := numberMap[digit]
		lettersCount := len(letters)
		for i := 0; i < lettersCount; i++ {
			backTrack(digits, index+1, combinations+string(letters[i]))
		}
	} else {
		retSli = append(retSli, combinations)
	}
}

// letterCombinationsQueue 队列解法。类似 BFS。
func letterCombinationsQueue(digits string) []string {
	if len(digits) == 0 {
		return make([]string, 0)
	}
	start := numberMap[string(digits[0])]
	queue := make([]string, 0, len(start))
	queue = append(queue, start...)
	for _, d := range digits[1:] {
		values := numberMap[string(d)]

		n := len(queue)
		// 模拟队列，取出上次循环的值
		for i := 0; i < n; i++ {
			// 模拟取出操作
			pop := queue[0]
			queue = queue[1:]
			// 依次添加
			for _, t := range values {
				queue = append(queue, pop+t)
			}
		}
		//// 如果没有元素，那么认为是第一次添加，此时需要
		//if len(queue) == 0 {
		//}
	}
	return queue
}
