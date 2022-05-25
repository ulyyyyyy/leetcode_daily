package p3

import (
	"fmt"
	"testing"
)

func Test_calcSlop(t *testing.T) {
	slop := calcSlop([]int{1, 7}, []int{2, 6})
	fmt.Println(slop)
}

func Test_minimumLines(t *testing.T) {
	fmt.Println(minimumLines([][]int{{1, 1000000000}, {1000000000, 1000000000}, {999999999, 1}, {2, 999999999}}))
}
