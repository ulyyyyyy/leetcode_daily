package ymd220127

import (
	"strings"
)

var (
	opExM  = int32(33) // !
	opSep  = int32(44) // ,
	opPt   = int32(46) // .
	opLine = int32(45) // -
)

func countValidWords(sentence string) (count int) {
	// 拿到每个词
	for _, word := range strings.Split(sentence, " ") {
		if checkWord(strings.TrimSpace(word)) {
			count++
		}
	}
	return count
}

func checkWord(word string) (rlt bool) {
	if len(word) == 0 {
		return false
	}
	isLine := false
	for index, ch := range word {
		length := len(word)

		// 标点符号
		if ch == opExM || ch == opSep || ch == opPt {
			// 如果最后一位是标点
			return index == length-1
		}

		if ch == opLine {
			if index >= length-1 || index == 0 {
				return false
			}
			prv := word[index-1]
			next := word[index+1]

			if !isLine && isLowerCase(prv) && isLowerCase(next) {
				isLine = true
				continue
			}
			return false
		}
		// 如果是小写
		if !isLowerCase(uint8(ch)) {
			return false
		}
	}
	return true
}

func isLowerCase(s uint8) (rlt bool) {
	return s >= 97 && s <= 122
}
