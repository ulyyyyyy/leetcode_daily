package d12

import "testing"

func Test_numEnclaves(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
	}{
		{
			name: "飞地的数量-测试用例1",
			args: args{
				grid: [][]int{{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}},
			},
			wantCount: 3,
		},
		{
			name: "飞地的数量-测试用例2",
			args: args{
				grid: [][]int{{0, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}},
			},
			wantCount: 0,
		},
		{
			name: "飞地的数量-测试用例3-err",
			args: args{
				grid: [][]int{
					{0, 0, 0, 1, 1, 1, 0, 1, 0, 0},
					{1, 1, 0, 0, 0, 1, 0, 1, 1, 1},
					{0, 0, 0, 1, 1, 1, 0, 1, 0, 0},
					{0, 1, 1, 0, 0, 0, 1, 0, 1, 0},
					{0, 1, 1, 1, 1, 1, 0, 0, 1, 0},
					{0, 0, 1, 0, 1, 1, 1, 1, 0, 1},
					{0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
					{0, 0, 1, 0, 0, 1, 0, 1, 0, 1},
					{1, 0, 1, 0, 1, 1, 0, 0, 0, 0},
					{0, 0, 0, 0, 1, 1, 0, 0, 0, 1},
				},
			},
			wantCount: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := numEnclaves(tt.args.grid); gotCount != tt.wantCount {
				t.Errorf("numEnclaves() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numEnclaves([][]int{
			{0, 0, 0, 1, 1, 1, 0, 1, 0, 0},
			{1, 1, 0, 0, 0, 1, 0, 1, 1, 1},
			{0, 0, 0, 1, 1, 1, 0, 1, 0, 0},
			{0, 1, 1, 0, 0, 0, 1, 0, 1, 0},
			{0, 1, 1, 1, 1, 1, 0, 0, 1, 0},
			{0, 0, 1, 0, 1, 1, 1, 1, 0, 1},
			{0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
			{0, 0, 1, 0, 0, 1, 0, 1, 0, 1},
			{1, 0, 1, 0, 1, 1, 0, 0, 0, 0},
			{0, 0, 0, 0, 1, 1, 0, 0, 0, 1},
		})
	}
}
