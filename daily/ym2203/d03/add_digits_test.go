package d03

import "testing"

func Test_addDigits(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name:    "各位相加-测试用例1",
			args:    args{num: 38},
			wantAns: 2,
		},
		{
			name:    "各位相加-测试用例2",
			args:    args{num: 0},
			wantAns: 0,
		},
		{
			name:    "各位相加-测试用例3",
			args:    args{num: 10},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := addDigits(tt.args.num); gotAns != tt.wantAns {
				t.Errorf("addDigits() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
