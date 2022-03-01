package d01

import "testing"

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Z 字形变换-测试用例1",
			args: args{
				s:       "PAYPALISHIRING",
				numRows: 3,
			},
			want: "PAHNAPLSIIGYIR",
		},
		{
			name: "Z 字形变换-测试用例2",
			args: args{
				s:       "PAYPALISHIRING",
				numRows: 4,
			},
			want: "PINALSIGYAHRPI",
		},
		{
			name: "Z 字形变换-测试用例3",
			args: args{
				s:       "A",
				numRows: 1,
			},
			want: "A",
		},
		{
			name: "Z 字形变换-测试用例4",
			args: args{
				s:       "AB",
				numRows: 1,
			},
			want: "AB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
