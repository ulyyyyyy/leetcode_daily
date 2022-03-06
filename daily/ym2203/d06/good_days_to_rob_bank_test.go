package d06

import (
	"reflect"
	"testing"
)

func Test_goodDaysToRobBank(t *testing.T) {
	type args struct {
		security []int
		time     int
	}
	tests := []struct {
		name         string
		args         args
		wantGoodDays []int
	}{
		{
			name: "适合打劫银行的日子-测试用例1",
			args: args{
				security: []int{5, 3, 3, 3, 5, 6, 2},
				time:     2,
			},
			wantGoodDays: []int{2, 3},
		},
		{
			name: "适合打劫银行的日子-测试用例2",
			args: args{
				security: []int{1, 1, 1, 1, 1},
				time:     0,
			},
			wantGoodDays: []int{0, 1, 2, 3, 4},
		},
		{
			name: "适合打劫银行的日子-测试用例3",
			args: args{
				security: []int{1},
				time:     5,
			},
			wantGoodDays: []int{},
		},
		{
			name: "适合打劫银行的日子-测试用例4-err",
			args: args{
				security: []int{4, 3, 2, 1},
				time:     1,
			},
			wantGoodDays: []int{},
		},
		{
			name: "适合打劫银行的日子-测试用例5-err",
			args: args{
				security: []int{3, 1, 20, 17, 26, 0, 2, 29, 24},
				time:     3,
			},
			wantGoodDays: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotGoodDays := goodDaysToRobBank(tt.args.security, tt.args.time); !reflect.DeepEqual(gotGoodDays, tt.wantGoodDays) {
				t.Errorf("goodDaysToRobBank() = %v, want %v", gotGoodDays, tt.wantGoodDays)
			}
		})
	}
}
