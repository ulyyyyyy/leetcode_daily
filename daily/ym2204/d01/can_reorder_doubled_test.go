package d01

import "testing"

func Test_canReorderDoubled(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "二倍数对数组-测试用例1",
			args: args{
				arr: []int{3, 1, 3, 6},
			},
			want: false,
		},
		{
			name: "二倍数对数组-测试用例2",
			args: args{
				arr: []int{2, 1, 2, 6},
			},
			want: false,
		},
		{
			name: "二倍数对数组-测试用例3",
			args: args{
				arr: []int{4, -2, 2, -4},
			},
			want: true,
		},
		{
			name: "二倍数对数组-测试用例4-err",
			args: args{
				arr: []int{4, 4, 8, 8, 8, 8, 2, 2, 16, 16, 1, 32},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canReorderDoubled(tt.args.arr); got != tt.want {
				t.Errorf("canReorderDoubled() = %v, want %v", got, tt.want)
			}
		})
	}
}
