package d01

import (
	"fmt"
	"testing"
)

func Test_longestNiceSubstring(t *testing.T) {
	cases := []string{
		"YazaAay",
		"Bb",
		"c",
		"dDzeE",
	}

	for _, c := range cases {
		fmt.Println(longestNiceSubstring(c))
	}
}
