package p238

import (
	"reflect"
	"testing"
)

func Test_productExceptSelf2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name       string
		args       args
		wantAnswer []int
	}{
		{
			name: "除自身以外数组的乘积-测试用例1",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			wantAnswer: []int{24, 12, 8, 6},
		},
		{
			name: "除自身以外数组的乘积-测试用例1",
			args: args{
				nums: []int{-1, 1, 0, -3, 3},
			},
			wantAnswer: []int{0, 0, 9, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswer := productExceptSelf2(tt.args.nums); !reflect.DeepEqual(gotAnswer, tt.wantAnswer) {
				t.Errorf("productExceptSelf2() = %v, want %v", gotAnswer, tt.wantAnswer)
			}
		})
	}
}
