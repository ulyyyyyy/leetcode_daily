package ymd220126

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
}

func TestDetectSquares_Add(t *testing.T) {
}

func TestDetectSquares_Count(t *testing.T) {
}

func TestDetectSquares_init(t *testing.T) {
	squares := Constructor()

	(&squares).Add([]int{5, 10})
	(&squares).Add([]int{10, 5})
	(&squares).Add([]int{10, 10})

	(&squares).Add([]int{3, 0})
	(&squares).Add([]int{8, 0})
	(&squares).Add([]int{8, 5})

	(&squares).Add([]int{9, 0})
	(&squares).Add([]int{9, 8})
	(&squares).Add([]int{1, 8})

	(&squares).Add([]int{0, 0})
	(&squares).Add([]int{8, 0})
	(&squares).Add([]int{8, 8})

	fmt.Println(squares.Count([]int{0, 8}))
}
