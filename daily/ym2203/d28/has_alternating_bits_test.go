package d28

import "testing"

func Test_hasAlternatingBits(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "交替位二进制数-测试用例1",
			args: args{
				n: 5,
			},
			want: true,
		},
		{
			name: "交替位二进制数-测试用例2",
			args: args{
				n: 7,
			},
			want: false,
		},
		{
			name: "交替位二进制数-测试用例3",
			args: args{
				n: 11,
			},
			want: false,
		},
		{
			name: "交替位二进制数-测试用例4-err",
			args: args{
				n: 3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasAlternatingBits2(tt.args.n); got != tt.want {
				t.Errorf("hasAlternatingBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
