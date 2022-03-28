package d27

import (
	"fmt"
	"testing"
)

func TestMissingRolls(t *testing.T) {
	ret1 := missingRolls([]int{3, 2, 4, 3}, 4, 2)
	fmt.Println(ret1)

	ret2 := missingRolls([]int{1, 5, 6}, 3, 4)
	fmt.Println(ret2)

	ret3 := missingRolls([]int{1, 2, 3, 4}, 6, 4)
	fmt.Println(ret3)

	ret4 := missingRolls([]int{6, 3, 4, 3, 5, 3}, 1, 6)
	fmt.Println(ret4)

	ret5 := missingRolls([]int{4, 5, 6, 2, 3, 6, 5, 4, 6, 4, 5, 1, 6, 3, 1, 4, 5, 5, 3, 2, 3, 5, 3, 2, 1, 5, 4, 3, 5, 1, 5}, 4, 40)
	fmt.Println(ret5)
}
