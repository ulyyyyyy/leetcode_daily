package d07

import "testing"

func Test_convertToBase(t *testing.T) {
	type args struct {
		num int
		r   int
	}
	tests := []struct {
		name    string
		args    args
		wantAns string
	}{
		{
			name: "R 进制转换器",
			args: args{
				num: 100,
				r:   2,
			},
			wantAns: "1100100",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := convertToBase(tt.args.num, tt.args.r); gotAns != tt.wantAns {
				t.Errorf("convertToBase() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
