package d24

import (
	"reflect"
	"testing"
)

func Test_findBall(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name        string
		args        args
		wantAnswers []int
	}{
		{
			name: "球会落何处-测试用例1",
			args: args{grid: [][]int{
				{1, 1, 1, -1, -1},
				{1, 1, 1, -1, -1},
				{-1, -1, -1, 1, 1},
				{1, 1, 1, 1, -1},
				{-1, -1, -1, -1, -1},
			}},
			wantAnswers: []int{1, -1, -1, -1, -1},
		},
		{
			name: "球会落何处-测试用例2",
			args: args{grid: [][]int{
				{-1},
			}},
			wantAnswers: []int{-1},
		},
		{
			name: "球会落何处-测试用例3",
			args: args{grid: [][]int{
				{1, 1, 1, 1, 1, 1},
				{-1, -1, -1, -1, -1, -1},
				{1, 1, 1, 1, 1, 1},
				{-1, -1, -1, -1, -1, -1},
			}},
			wantAnswers: []int{0, 1, 2, 3, 4, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAnswers := findBall(tt.args.grid); !reflect.DeepEqual(gotAnswers, tt.wantAnswers) {
				t.Errorf("findBall() = %v, want %v", gotAnswers, tt.wantAnswers)
			}
		})
	}
}
