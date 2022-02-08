package ymd220203

import "testing"

func Test_findMinFibonacciNumbers(t *testing.T) {
	type args struct {
		k int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "和为 K 的最少斐波那契数字数目-测试用例1",
			args:    args{k: 7},
			wantAns: 2,
		},
		{
			name:    "和为 K 的最少斐波那契数字数目-测试用例2",
			args:    args{k: 10},
			wantAns: 2,
		},
		{
			name:    "和为 K 的最少斐波那契数字数目-测试用例3",
			args:    args{k: 19},
			wantAns: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findMinFibonacciNumbers(tt.args.k); gotAns != tt.wantAns {
				t.Errorf("findMinFibonacciNumbers() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
