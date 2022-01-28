package ymd220128

import (
	"fmt"
	"testing"
)

func Test_numberOfWeakCharacters(t *testing.T) {
	cases := [][][]int{
		{{5, 5}, {6, 3}, {3, 6}},
		{{2, 2}, {3, 3}},
		{{1, 5}, {10, 4}, {4, 3}},
	}
	for _, c := range cases {
		characters := numberOfWeakCharacters(c)
		fmt.Println(characters)
	}
}
