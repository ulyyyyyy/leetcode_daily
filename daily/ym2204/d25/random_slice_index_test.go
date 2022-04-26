package d25

import (
	"fmt"
	"testing"
)

func Test_getRandomEle(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(getRandomEle([]int{1, 2, 3, 4, 5}))
	}
}
