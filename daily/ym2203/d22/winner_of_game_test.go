package d22

import "testing"

func Test_winnerOfGame(t *testing.T) {
	type args struct {
		colors string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "如果相邻两个颜色均相同则删除当前颜色-测试用例1",
			args: args{
				colors: "AAABABB",
			},
			want: true,
		},
		{
			name: "如果相邻两个颜色均相同则删除当前颜色-测试用例2",
			args: args{
				colors: "AA",
			},
			want: false,
		},
		{
			name: "如果相邻两个颜色均相同则删除当前颜色-测试用例3",
			args: args{
				colors: "ABBBBBBBAAA",
			},
			want: false,
		},
		{
			name: "如果相邻两个颜色均相同则删除当前颜色-测试用例4",
			args: args{
				colors: "AAAAAA",
			},
			want: true,
		},
		{
			name: "如果相邻两个颜色均相同则删除当前颜色-测试用例5",
			args: args{
				colors: "B",
			},
			want: false,
		},
		{
			name: "如果相邻两个颜色均相同则删除当前颜色-测试用例6",
			args: args{
				colors: "ABABABABABABABABAAABAB",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := winnerOfGame(tt.args.colors); got != tt.want {
				t.Errorf("winnerOfGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_winnerOfGame2(t *testing.T) {
	type args struct {
		colors string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "如果相邻两个颜色均相同则删除当前颜色-测试用例1",
			args: args{
				colors: "AAABABB",
			},
			want: true,
		},
		{
			name: "如果相邻两个颜色均相同则删除当前颜色-测试用例2",
			args: args{
				colors: "AA",
			},
			want: false,
		},
		{
			name: "如果相邻两个颜色均相同则删除当前颜色-测试用例3",
			args: args{
				colors: "ABBBBBBBAAA",
			},
			want: false,
		},
		{
			name: "如果相邻两个颜色均相同则删除当前颜色-测试用例4-err",
			args: args{
				colors: "BBBBAAAA",
			},
			want: false,
		},
		{
			name: "如果相邻两个颜色均相同则删除当前颜色-测试用例5-err",
			args: args{
				colors: "AABB",
			},
			want: false,
		},
		{
			name: "如果相邻两个颜色均相同则删除当前颜色-测试用例6-err",
			args: args{
				colors: "AAAABBBB",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := winnerOfGame2(tt.args.colors); got != tt.want {
				t.Errorf("winnerOfGame2() = %v, want %v", got, tt.want)
			}
		})
	}
}
