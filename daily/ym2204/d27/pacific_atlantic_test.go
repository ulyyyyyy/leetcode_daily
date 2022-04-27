package d27

import (
	"reflect"
	"testing"
)

func Test_pacificAtlantic(t *testing.T) {
	type args struct {
		heights [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns [][]int
	}{
		{
			name: "太平洋大西洋水流问题-测试用例1",
			args: args{
				heights: [][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}},
			},
			wantAns: [][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}},
		},
		{
			name: "太平洋大西洋水流问题-测试用例2",
			args: args{
				heights: [][]int{{2, 1}, {1, 2}},
			},
			wantAns: [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := pacificAtlantic(tt.args.heights); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("pacificAtlantic() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

func BenchmarkPacificAtlantic1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pacificAtlantic([][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}})
	}
}

func BenchmarkPacificAtlantic2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pacificAtlantic2([][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}})
	}
}
