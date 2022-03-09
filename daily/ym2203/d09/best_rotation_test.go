package d09

import "testing"

func Test_bestRotation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "得分最高的最小轮调-测试用例1",
			args: args{
				nums: []int{2, 3, 1, 4, 0},
			},
			want: 3,
		},
		{
			name: "得分最高的最小轮调-测试用例2",
			args: args{
				nums: []int{1, 3, 0, 2, 4},
			},
			want: 0,
		},
		{
			name: "得分最高的最小轮调-测试用例3-err",
			args: args{
				nums: []int{6, 2, 8, 3, 5, 2, 4, 3, 7, 6},
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bestRotation(tt.args.nums); got != tt.want {
				t.Errorf("bestRotation() = %v, want %v", got, tt.want)
			}
		})
	}
}
