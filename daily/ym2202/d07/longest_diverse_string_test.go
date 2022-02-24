package d07

import (
	"fmt"
	"testing"
)

func Test_longestDiverseString(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "最长快乐字符串-测试用例1",
			args: args{
				a: 1,
				b: 1,
				c: 7,
			},
			want: "",
		},
		{
			name: "最长快乐字符串-测试用例2",
			args: args{
				a: 2, b: 2, c: 7,
			},
			want: "",
		},
		{
			name: "最长快乐字符串-测试用例3",
			args: args{
				a: 7, b: 2, c: 1,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestDiverseString(tt.args.a, tt.args.b, tt.args.c)
			fmt.Println(got)
		})
	}
}
