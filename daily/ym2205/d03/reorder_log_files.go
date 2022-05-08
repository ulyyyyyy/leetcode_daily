package d03

import (
	"reflect"
	"sort"
	"strings"
	"unicode"
)

func reorderLogFiles(logs []string) []string {
	letterSli := make([]string, 0)
	digitSli := make([]string, 0)

	for _, log := range logs {
		// 判断是字母日志还是数字日志
		if isDigit(log) {
			digitSli = append(digitSli, log)
		} else {
			letterSli = append(letterSli, log)
		}
	}

	sort.Slice(letterSli, func(i, j int) bool {
		pre := strings.Split(letterSli[i], " ")
		last := strings.Split(letterSli[j], " ")

		preSynx, preLogs := pre[0], pre[1:]
		lastSynx, lastLogs := last[0], last[1:]

		//
		if reflect.DeepEqual(preLogs, lastLogs) {
			return sort.StringsAreSorted([]string{preSynx, lastSynx})
		}
		return sort.StringsAreSorted([]string{strings.Join(preLogs, " "), strings.Join(lastLogs, " ")})
	})
	letterSli = append(letterSli, digitSli...)
	return letterSli
}

func isDigit(log string) bool {
	spl := strings.Split(log, " ")
	first := spl[1][0]
	return !unicode.IsLetter(rune(first))
}
