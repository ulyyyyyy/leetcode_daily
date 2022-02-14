package ymd220212

import "testing"

func Test_numEnclavesOffice(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "飞地的数量-测试用例1",
			args: args{
				grid: [][]int{{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}},
			},
			wantAns: 3,
		},
		{
			name: "飞地的数量-测试用例2",
			args: args{
				grid: [][]int{{0, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}},
			},
			wantAns: 0,
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
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := numEnclavesOffice(tt.args.grid); gotAns != tt.wantAns {
				t.Errorf("numEnclavesOffice() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

func BenchmarkNumEnclavesOffice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		numEnclavesOffice([][]int{
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
