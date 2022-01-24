package ymd220124

import (
	"fmt"
	"testing"
)

func Test_calcWaitTime(t *testing.T) {
	fmt.Println(calcWaitTime(30, 4))
}

func Test_secondMinimum(t *testing.T) {
	//fmt.Println(secondMinimum(5, [][]int{{1, 2}, {1, 3}, {1, 4}, {3, 4}, {4, 5}}, 3, 5))
	fmt.Println(secondMinimum(2, [][]int{{1, 2}}, 1, 2))
}

func Test_officialEx(t *testing.T) {
	fmt.Println(officialEx(2, [][]int{{1, 2}}, 1, 2))
}
