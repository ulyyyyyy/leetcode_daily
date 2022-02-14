package ymd220130

import (
	"fmt"
	"testing"
)

func Test_uncommonFromSentences(t *testing.T) {

	testCases := [][]string{
		{"this apple is sweet", "this apple is sour"},
		{"apple apple", "banana"},
	}

	for _, testCase := range testCases {
		fmt.Println(uncommonFromSentences(testCase[0], testCase[1]))
	}
}
