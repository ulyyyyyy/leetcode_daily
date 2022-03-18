package d18

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	bank := Constructor([]int64{10, 100, 20, 50, 30})
	rltSlice := make([]bool, 0)

	rltSlice = append(rltSlice, bank.Withdraw(3, 10))
	rltSlice = append(rltSlice, bank.Transfer(5, 1, 20))
	rltSlice = append(rltSlice, bank.Deposit(5, 20))
	rltSlice = append(rltSlice, bank.Transfer(3, 4, 15))
	rltSlice = append(rltSlice, bank.Withdraw(10, 50))

	wantSlice := []bool{true, true, true, false, false}
	fmt.Println(reflect.DeepEqual(rltSlice, wantSlice))
}
