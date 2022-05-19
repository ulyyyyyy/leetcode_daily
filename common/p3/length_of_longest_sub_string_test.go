package p3

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "无重复字符的最长子串-测试用例1",
			args: args{
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "无重复字符的最长子串-测试用例2",
			args: args{
				s: "bbbbb",
			},
			want: 1,
		},
		{
			name: "无重复字符的最长子串-测试用例3",
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
