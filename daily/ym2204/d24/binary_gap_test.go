package d24

import "testing"

func Test_binaryGap(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "二进制间距-测试用例1",
			args: args{
				n: 22,
			},
			want: 2,
		},
		{
			name: "二进制间距-测试用例2",
			args: args{
				n: 8,
			},
			want: 0,
		},
		{
			name: "二进制间距-测试用例3",
			args: args{
				n: 5,
			},
			want: 2,
		},
		{
			name: "旋转函数-测试用例3",
			args: args{
				n: 25,
			},
			want: 3,
		},
		{
			name: "旋转函数-测试用例4",
			args: args{
				n: 764,
			},
			want: 2,
		},
		{
			name: "旋转函数-测试用例5",
			args: args{
				n: 426345,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binaryGap(tt.args.n); got != tt.want {
				t.Errorf("binaryGap() = %v, want %v", got, tt.want)
			}
		})
	}
}
