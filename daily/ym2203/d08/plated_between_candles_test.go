package d08

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_findPlates(t *testing.T) {
	fmt.Println(findPlates("|****|"))
	fmt.Println(findPlates("|***|*|"))
	fmt.Println(findPlates("|**|**|"))
	fmt.Println(findPlates("**|****|"))
	fmt.Println(findPlates("**|****|**"))
}

func Test_platesBetweenCandles(t *testing.T) {
	type args struct {
		s       string
		queries [][]int
	}
	tests := []struct {
		name        string
		args        args
		wantAnswers []int
	}{
		{
			name: "蜡烛之间的盘子-测试用例1",
			args: args{
				s:       "**|**|***|",
				queries: [][]int{{2, 5}, {5, 9}},
			},
			wantAnswers: []int{2, 3},
		},
		{
			name: "蜡烛之间的盘子-测试用例2",
			args: args{
				s:       "***|**|*****|**||**|*",
				queries: [][]int{{1, 17}, {4, 5}, {14, 17}, {5, 11}, {15, 16}},
			},
			wantAnswers: []int{9, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswers := platesBetweenCandlesOffice(tt.args.s, tt.args.queries); !reflect.DeepEqual(gotAnswers, tt.wantAnswers) {
				t.Errorf("platesBetweenCandles() = %v, want %v", gotAnswers, tt.wantAnswers)
			}
		})
	}
}
