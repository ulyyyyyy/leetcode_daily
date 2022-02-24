package d04

import "testing"

func Test_countGoodRectangles(t *testing.T) {
	type args struct {
		rectangles [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "可以形成最大正方形的矩形数目-测试用例1",
			args: args{
				rectangles: [][]int{{5, 8}, {3, 9}, {5, 12}, {16, 5}},
			},
			wantAns: 3,
		},
		{
			name: "可以形成最大正方形的矩形数目-测试用例2",
			args: args{
				rectangles: [][]int{{2, 3}, {3, 7}, {4, 3}, {3, 7}},
			},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countGoodRectangles(tt.args.rectangles); gotAns != tt.wantAns {
				t.Errorf("countGoodRectangles() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
