package d26

import "strconv"

// 你现在是一场采用特殊赛制棒球比赛的记录员。这场比赛由若干回合组成，过去几回合的得分可能会影响以后几回合的得分。
//
// 比赛开始时，记录是空白的。你会得到一个记录操作的字符串列表 ops，其中 ops[i] 是你需要记录的第 i 项操作，ops 遵循下述规则：
//
// 整数 x - 表示本回合新获得分数 x
// "+" - 表示本回合新获得的得分是前两次得分的总和。题目数据保证记录此操作时前面总是存在两个有效的分数。
// "D" - 表示本回合新获得的得分是前一次得分的两倍。题目数据保证记录此操作时前面总是存在一个有效的分数。
// "C" - 表示前一次得分无效，将其从记录中移除。题目数据保证记录此操作时前面总是存在一个有效的分数。
// 请你返回记录中所有得分的总和。

func calPoints(ops []string) int {
	total := make([]int, 0)
	for _, s := range ops {
		if s == "C" {
			total = total[:len(total)-1]
		} else if s == "D" {
			total = append(total, total[len(total)-1]*2)
		} else if s == "+" {
			total = append(total, total[len(total)-1]+total[len(total)-2])
		} else {
			v, _ := strconv.Atoi(s)
			total = append(total, v)
		}
	}
	ret := 0
	for _, t := range total {
		ret += t
	}
	return ret
}

func calPointsSwitch(ops []string) int {
	total := make([]int, 0)
	for _, s := range ops {
		switch s {
		case "C":
			total = total[:len(total)-1]
		case "D":
			total = append(total, total[len(total)-1]*2)
		case "+":
			total = append(total, total[len(total)-1]+total[len(total)-2])
		default:
			v, _ := strconv.Atoi(s)
			total = append(total, v)

		}
	}
	ret := 0
	for _, t := range total {
		ret += t
	}
	return ret
}
