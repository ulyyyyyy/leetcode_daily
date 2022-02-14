package ymd220127

import (
	"fmt"
	"testing"
)

func Test_countValidWords(t *testing.T) {
	fmt.Println(countValidWords("cat and  dog"))
	fmt.Println(countValidWords("!this  1-s b8d!"))
	fmt.Println(countValidWords("alice and  bob are playing stone-game10"))

	fmt.Println(countValidWords("he bought 2 pencils, 3 erasers, and 1  pencil-sharpener."))

	fmt.Println(countValidWords("-"))

	fmt.Println(countValidWords(". ! 7hk  al6 l! aon49esj35la k3 7u2tkh  7i9y5  !jyylhppd et v- h!ogsouv 5"))
}
