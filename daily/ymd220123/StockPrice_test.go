package ymd220123

import (
	"fmt"
	"testing"
)

func TestStockPrice_Minimum(t *testing.T) {
	sp := Constructor()
	sp.Update(38, 1)
	fmt.Println(sp.Current())
	fmt.Println(sp.Minimum())
	fmt.Println(sp.Maximum())
	sp.Update(39, 2)
	fmt.Println(sp.Current())
	fmt.Println(sp.Minimum())
	fmt.Println(sp.Maximum())
	sp.Update(40, 3)
	fmt.Println(sp.Current())
	fmt.Println(sp.Minimum())
	fmt.Println(sp.Maximum())
	sp.Update(38, 4)
	fmt.Println(sp.Current())
	fmt.Println(sp.Minimum())
	fmt.Println(sp.Maximum())
	sp.Update(40, 5)
	fmt.Println(sp.Current())
	fmt.Println(sp.Minimum())
	fmt.Println(sp.Maximum())
}
