package ymd220129

import (
	"testing"
)

func Test_highestPeak(t *testing.T) {
	cases := [][][]int{
		{{0, 1}, {0, 0}},
		{{0, 0, 1}, {1, 0, 0}, {0, 0, 0}},
	}

	for _, c := range cases {
		highestPeak(c)
	}
}
