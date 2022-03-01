package d26

import "testing"

func Test_maximumDifference(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name    string
		args    args
		wantMin int
	}{
		{
			name:    "增量元素之间的最大差值-测试用例1",
			args:    args{[]int{7, 1, 5, 4}},
			wantMin: 4,
		},
		{
			name:    "增量元素之间的最大差值-测试用例2",
			args:    args{[]int{7, 1, 5, 4}},
			wantMin: 4,
		},
		{
			name:    "增量元素之间的最大差值-测试用例3",
			args:    args{[]int{7, 1, 5, 4}},
			wantMin: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMin := maximumDifference(tt.args.nums); gotMin != tt.wantMin {
				t.Errorf("maximumDifference() = %v, want %v", gotMin, tt.wantMin)
			}
		})
	}
}
