package d26

import "testing"

func Test_projectionArea(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "三维形体投影面积-测试用例1",
			args: args{
				grid: [][]int{{1, 2}, {3, 4}},
			},
			want: 17,
		},
		{
			name: "三维形体投影面积-测试用例2",
			args: args{
				grid: [][]int{{2}},
			},
			want: 5,
		},
		{
			name: "三维形体投影面积-测试用例3",
			args: args{
				grid: [][]int{{1, 0}, {0, 2}},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := projectionArea(tt.args.grid); got != tt.want {
				t.Errorf("projectionArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
