package ymd220122

import (
	"fmt"
	"testing"
)

var (
	testcase = []string{
		"",
		"abba",
		"abbbaa",
		"abaabb",
		"abbbbbbbbbaabbaba",
	}
)

func Test_removePalindromeSub(t *testing.T) {
	for _, c := range testcase {
		count := removePalindromeSub(c)
		fmt.Println(count)
	}
}
