package d02

import (
	"fmt"
	"testing"
)

func Test_reversePrefix(t *testing.T) {
	type i struct {
		word string
		ch   byte
	}
	cases := []i{
		{word: "abcdefd", ch: 'd'},
		{word: "xyxzxe", ch: 'z'},
		{word: "abcd", ch: 'z'},
	}

	for _, c := range cases {
		fmt.Println(reversePrefix(c.word, c.ch))
	}
}
