package p31

import "testing"

func Test_nextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "下一个排列-测试用例1",
			args: args{
				nums: []int{1, 2, 3},
			},
		},
		{
			name: "下一个排列-测试用例2",
			args: args{
				nums: []int{3, 2, 1},
			},
		},
		{
			name: "下一个排列-测试用例3-err",
			args: args{
				nums: []int{1, 3, 2},
				// 2, 1, 3
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.args.nums)
		})
	}
}
