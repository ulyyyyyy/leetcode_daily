package d25

import "testing"

func Test_trailingZeroes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantCnt int
	}{
		{
			name: "阶乘后的零-测试用例1",
			args: args{
				n: 3,
			},
			wantCnt: 0,
		},
		{
			name: "阶乘后的零-测试用例2",
			args: args{
				n: 5,
			},
			wantCnt: 1,
		},
		{
			name: "阶乘后的零-测试用例3",
			args: args{
				n: 0,
			},
			wantCnt: 0,
		},
		{
			name: "阶乘后的零-测试用例4-err",
			args: args{
				n: 30,
			},
			wantCnt: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCnt := trailingZeroes2(tt.args.n); gotCnt != tt.wantCnt {
				t.Errorf("trailingZeroes() = %v, want %v", gotCnt, tt.wantCnt)
			}
		})
	}
}
