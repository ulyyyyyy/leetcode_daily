package ymd220214

import "testing"

func Test_singleNonDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "有序数组中的单一元素-测试用例1",
			args: args{
				nums: []int{1, 1, 2, 3, 3, 4, 4, 8, 8},
			},
			wantAns: 2,
		},
		{
			name: "有序数组中的单一元素-测试用例2",
			args: args{
				nums: []int{3, 3, 7, 7, 10, 11, 11},
			},
			wantAns: 10,
		},
		{
			name: "有序数组中的单一元素-测试用例3",
			args: args{
				nums: []int{1, 2, 2},
			},
			wantAns: 1,
		},
		{
			name: "有序数组中的单一元素-测试用例4",
			args: args{
				nums: []int{3, 3, 4, 4, 10, 10, 15, 15, 20, 20, 1111, 1111, 2222},
			},
			wantAns: 2222,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := singleNonDuplicate(tt.args.nums); gotAns != tt.wantAns {
				t.Errorf("singleNonDuplicate() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
