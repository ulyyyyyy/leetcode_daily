package prefix_count

import "strings"

func prefixCount(words []string, pref string) (count int) {
	for _, word := range words {
		if strings.HasPrefix(word, pref) {
			count++
		}
	}
	return
}
