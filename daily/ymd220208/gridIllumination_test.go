package ymd220208

import (
	"reflect"
	"testing"
)

func Test_gridIllumination(t *testing.T) {
	type args struct {
		n       int
		lamps   [][]int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "网格照明-测试用例1",
			args: args{
				n:       5,
				lamps:   [][]int{{0, 0}, {4, 4}},
				queries: [][]int{{1, 1}, {1, 0}},
			},
			want: []int{1, 0},
		},
		{
			name: "网格照明-测试用例2",
			args: args{
				n:       5,
				lamps:   [][]int{{0, 0}, {4, 4}},
				queries: [][]int{{1, 1}, {1, 1}},
			},
			want: []int{1, 1},
		},
		{
			name: "网格照明-测试用例3",
			args: args{
				n:       5,
				lamps:   [][]int{{0, 0}, {0, 4}},
				queries: [][]int{{0, 4}, {0, 1}, {1, 4}},
			},
			want: []int{1, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gridIllumination(tt.args.n, tt.args.lamps, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("gridIllumination() = %v, want %v", got, tt.want)
			}
		})
	}
}
