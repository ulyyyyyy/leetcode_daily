package ymd220206

import "testing"

func Test_sumOfUnique(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
	}{
		{
			name: "唯一元素的和-测试用例1",
			args: args{
				nums: []int{1, 2, 3, 2},
			},
			wantCount: 4,
		},
		{
			name: "唯一元素的和-测试用例2",
			args: args{
				nums: []int{1, 1, 1, 1, 1},
			},
			wantCount: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := sumOfUnique(tt.args.nums); gotCount != tt.wantCount {
				t.Errorf("sumOfUnique() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
