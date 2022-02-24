package d13

import "testing"

func Test_maxNumberOfBalloons(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: " “气球” 的最大数量-测试用例1",
			args: args{
				text: "nlaebolko",
			},
			wantAns: 1,
		},
		{
			name: " “气球” 的最大数量-测试用例2",
			args: args{
				text: "loonbalxballpoon",
			},
			wantAns: 2,
		},
		{
			name: " “气球” 的最大数量-测试用例3",
			args: args{
				text: "leetcode",
			},
			wantAns: 0,
		},
		{
			name: " “气球” 的最大数量-测试用例4-err",
			args: args{
				text: "balon",
			},
			wantAns: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := maxNumberOfBalloons(tt.args.text); gotAns != tt.wantAns {
				t.Errorf("maxNumberOfBalloons() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
