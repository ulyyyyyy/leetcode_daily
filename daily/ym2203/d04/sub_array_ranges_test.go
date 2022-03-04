package d04

import "testing"

func Test_subArrayRanges(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int64
	}{
		{
			name:    "子数组范围和-测试用例1",
			args:    args{nums: []int{1, 2, 3}},
			wantAns: 4,
		},
		{
			name:    "子数组范围和-测试用例2",
			args:    args{nums: []int{1, 3, 3}},
			wantAns: 4,
		},
		{
			name:    "子数组范围和-测试用例3",
			args:    args{nums: []int{4, -2, -3, 4, 1}},
			wantAns: 59,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := subArrayRanges(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("subArrayRanges() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
