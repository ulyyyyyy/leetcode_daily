package ymd220121

import "testing"

var (
	arr = []int{100, -23, -23, 404, 100, 23, 23, 23, 3, 404}
)

func Test_generaMap(t *testing.T) {
	generaMap()
}

func Test_minJumps(t *testing.T) {
	minJumps(arr)
}
