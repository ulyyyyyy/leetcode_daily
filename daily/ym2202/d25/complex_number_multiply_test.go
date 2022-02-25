package d25

import "testing"

// BenchmarkConv
// BenchmarkConv-16     	154338453	         7.671 ns/op
// BenchmarkConv2
// BenchmarkConv2-16    	17630983	        66.66 ns/op
// 两者时间消耗差异巨大

func BenchmarkConv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conv("1+1i")
	}
}

func BenchmarkConv2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conv2("1+1i")
	}
}

func Test_complexNumberMultiply(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "复数乘法-测试用例1",
			args: args{
				num1: "1+1i",
				num2: "1+1i",
			},
			want: "0+2i",
		},
		{
			name: "复数乘法-测试用例2",
			args: args{
				num1: "1+-1i",
				num2: "1+-1i",
			},
			want: "0+-2i",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := complexNumberMultiply(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("complexNumberMultiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
