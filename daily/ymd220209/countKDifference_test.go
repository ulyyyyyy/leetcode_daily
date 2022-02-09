package ymd220209

import "testing"

func Test_countKDifference(t *testing.T) {
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
			name: "差的绝对值为 K 的数对数目-测试用例1",
			args: args{
				nums: []int{1, 2, 2, 1},
				k:    1,
			},
			wantAns: 4,
		},
		{
			name: "差的绝对值为 K 的数对数目-测试用例2",
			args: args{
				nums: []int{1, 3},
				k:    3,
			},
			wantAns: 0,
		},
		{
			name: "差的绝对值为 K 的数对数目-测试用例3",
			args: args{
				nums: []int{3, 2, 1, 5, 4},
				k:    2,
			},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countKDifference(tt.args.nums, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("countKDifference() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
