package d20

import (
	"reflect"
	"testing"
)

func Test_findRightInterval(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "寻找右区间-测试用例1",
			args: args{
				intervals: [][]int{
					{1, 2},
				},
			},
			wantAns: []int{-1},
		},
		{
			name: "寻找右区间-测试用例2",
			args: args{
				intervals: [][]int{
					{3, 4},
					{2, 3},
					{1, 2},
				},
			},
			wantAns: []int{-1, 0, 1},
		},
		{
			name: "寻找右区间-测试用例3",
			args: args{
				intervals: [][]int{
					{1, 4},
					{2, 3},
					{3, 4},
				},
			},
			wantAns: []int{-1, 2, -1},
		},
		{
			name: "寻找右区间-测试用例4-err",
			args: args{
				intervals: [][]int{
					{1, 2},
					{2, 3},
					{0, 1},
					{3, 4},
				},
			},
			wantAns: []int{1, 3, 0, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findRightInterval(tt.args.intervals); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("findRightInterval() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

func Test_binarySearch(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int
	}{
		{
			name: "",
			args: args{
				nums:   []int{1, 2, 3, 4, 5, 6},
				target: 1,
			},
			wantRes: 0,
		},
		{
			name: "",
			args: args{
				nums:   []int{1, 2, 3, 4, 5, 6},
				target: 6,
			},
			wantRes: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := binarySearch(tt.args.nums, tt.args.target); gotRes != tt.wantRes {
				t.Errorf("binarySearch() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
