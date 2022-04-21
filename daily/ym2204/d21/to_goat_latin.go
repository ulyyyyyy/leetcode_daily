package d21

import "strings"

var vowelMap = map[byte]struct{}{
	'a': {},
	'e': {},
	'i': {},
	'o': {},
	'u': {},
	'A': {},
	'E': {},
	'I': {},
	'O': {},
	'U': {},
}

func toGoatLatin(sentence string) string {
	//
	tmpSli := make([]string, 0)
	wordNum := 1

	tmpWord := ""
	for _, r := range sentence {
		if byte(r) != ' ' {
			tmpWord += string(r)
			continue
		}
		tmpSli = append(tmpSli, wordHandler(tmpWord, wordNum))
		tmpWord = ""
		wordNum++
	}
	if tmpWord != "" {
		tmpSli = append(tmpSli, wordHandler(tmpWord, wordNum))
	}
	return strings.Join(tmpSli, " ")
}

func wordHandler(word string, wordNum int) string {
	// 如果该字符不是元音，那么调整
	if _, ok := vowelMap[word[0]]; !ok {
		word = word[1:] + string(word[0])
	}
	word += "ma"
	for i := 0; i < wordNum; i++ {
		word += "a"
	}
	return word
}
