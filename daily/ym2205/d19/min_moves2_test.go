package d19

import "testing"

func Test_minMoves2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "最少移动次数使数组元素相等 II-测试用例1",
			args: args{
				nums: []int{1, 2, 3},
			},
			wantAns: 2,
		},
		{
			name: "最少移动次数使数组元素相等 II-测试用例2",
			args: args{
				nums: []int{1, 10, 2, 9},
			},
			wantAns: 16,
		},
		{
			name: "最少移动次数使数组元素相等 II-测试用例3",
			args: args{
				nums: []int{1, 2, 3, 4, 5},
			},
			wantAns: 6,
		},
		{
			name: "最少移动次数使数组元素相等 II-测试用例4-err",
			args: args{
				nums: []int{12345, 2345, 36541, 23456, 76574, 54362, 12347, 99999, 77675, 3452, 33},
			},
			wantAns: 314629,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minMoves2(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("minMoves2() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
