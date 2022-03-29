package d29

import "testing"

func Test_maxConsecutiveAnswers(t *testing.T) {
	type args struct {
		answerKey string
		k         int
	}
	tests := []struct {
		name    string
		args    args
		wantMax int
	}{
		{
			name: "考试的最大困扰度-测试用例1",
			args: args{
				answerKey: "TTFF",
				k:         2,
			},
			wantMax: 4,
		},
		{
			name: "考试的最大困扰度-测试用例2",
			args: args{
				answerKey: "TFFT",
				k:         1,
			},
			wantMax: 3,
		},
		{
			name: "考试的最大困扰度-测试用例3",
			args: args{
				answerKey: "TTFTTFTT",
				k:         1,
			},
			wantMax: 5,
		},
		{
			name: "考试的最大困扰度-测试用例4",
			args: args{
				answerKey: "TTTTTT",
				k:         1,
			},
			wantMax: 6,
		},
		{
			name: "考试的最大困扰度-测试用例5",
			args: args{
				answerKey: "TFTFTFTFTFTFTF",
				k:         1,
			},
			wantMax: 3,
		},
		{
			name: "考试的最大困扰度-测试用例6",
			args: args{
				answerKey: "TFTFTFTFTFTFTF",
				k:         111,
			},
			wantMax: 14,
		},
		{
			name: "考试的最大困扰度-测试用例7",
			args: args{
				answerKey: "TFTFTFTFTFTFTF",
				k:         111,
			},
			wantMax: 14,
		},
		{
			name: "考试的最大困扰度-测试用例8",
			args: args{
				answerKey: "TFTFTFTFTFTFTF",
				k:         4,
			},
			wantMax: 9,
		},
		{
			name: "考试的最大困扰度-测试用例9",
			args: args{
				answerKey: "FFFTTFTTFT",
				k:         3,
			},
			wantMax: 8,
		},
		{
			name: "考试的最大困扰度-测试用例10",
			args: args{
				answerKey: "TTFTTTTTFT",
				k:         1,
			},
			wantMax: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMax := maxConsecutiveAnswers(tt.args.answerKey, tt.args.k); gotMax != tt.wantMax {
				t.Errorf("maxConsecutiveAnswers() = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}
