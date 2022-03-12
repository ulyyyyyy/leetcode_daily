package p55

import "testing"

func Test_canJump(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "跳跃游戏-测试用例1",
			args: args{[]int{2, 3, 1, 1, 4}},
			want: true,
		},
		{
			name: "跳跃游戏-测试用例2",
			args: args{[]int{3, 2, 1, 0, 4}},
			want: false,
		},
		{
			name: "跳跃游戏-测试用例3-err",
			args: args{[]int{0}},
			want: true,
		},
		{
			name: "跳跃游戏-测试用例4-err",
			args: args{[]int{0, 1}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.args.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
