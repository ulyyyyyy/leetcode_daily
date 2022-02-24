package d08

import (
	"fmt"
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

		{
			name: "网格照明-测试用例4-err",
			args: args{
				n:       6,
				lamps:   [][]int{{2, 5}, {4, 2}, {0, 3}, {0, 5}, {1, 4}, {4, 2}, {3, 3}, {1, 0}},
				queries: [][]int{{4, 3}, {3, 1}, {5, 3}, {0, 5}, {4, 4}, {3, 3}},
			},
			want: []int{1, 0, 1, 1, 0, 1},
		},

		{
			name: "网格照明-测试用例5-err",
			args: args{
				n:       5,
				lamps:   [][]int{{0, 0}, {4, 4}},
				queries: [][]int{{1, 1}, {0, 0}, {2, 2}},
			},
			want: []int{1, 1, 1},
		},

		{
			name: "网格照明-测试用例6-err",
			args: args{
				n:       5,
				lamps:   [][]int{{0, 0}, {0, 0}, {4, 4}},
				queries: [][]int{{1, 1}, {1, 0}},
			},
			want: []int{1, 0},
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

func TestName(t *testing.T) {
	b := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(b); i++ {
		fmt.Println(i, b[i])
		if b[i] == 3 {
			b = append(b[:i], b[i+1:]...)
			i--
		}
	}
}
