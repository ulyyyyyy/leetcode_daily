package d11

import "testing"

func Test_minimumDifference(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "学生分数最小差值-测试用例1",
			args: args{
				nums: []int{90},
				k:    1,
			},
			wantAns: 0,
		},
		{
			name: "学生分数最小差值-测试用例2",
			args: args{
				nums: []int{9, 4, 1, 7},
				k:    2,
			},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minimumDifference(tt.args.nums, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("minimumDifference() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
