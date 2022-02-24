package d15

import (
	"reflect"
	"testing"
)

func Test_luckyNumbers(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name      string
		args      args
		wantSlice []int
	}{
		{
			name: "矩阵中的幸运数-测试用例1",
			args: args{
				matrix: [][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}},
			},
			wantSlice: []int{15},
		},
		{
			name: "矩阵中的幸运数-测试用例2",
			args: args{
				matrix: [][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}},
			},
			wantSlice: []int{15},
		},
		{
			name: "矩阵中的幸运数-测试用例3",
			args: args{
				matrix: [][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}},
			},
			wantSlice: []int{15},
		},
		{
			name: "矩阵中的幸运数-测试用例4-err",
			args: args{
				matrix: [][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}},
			},
			wantSlice: []int{12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSlice := luckyNumbers(tt.args.matrix); !reflect.DeepEqual(gotSlice, tt.wantSlice) {
				t.Errorf("luckyNumbers() = %v, want %v", gotSlice, tt.wantSlice)
			}
		})
	}
}
