package ymd220131

import (
	"fmt"
	"testing"
)

func Test_numberOfSteps(t *testing.T) {
	cases := []int{
		14,
		8,
		123,
	}

	for _, c := range cases {
		fmt.Println(numberOfSteps(c))
	}
}
