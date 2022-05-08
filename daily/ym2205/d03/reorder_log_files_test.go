package d03

import (
	"reflect"
	"testing"
)

func Test_isDigit(t *testing.T) {
	type args struct {
		log string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "是否是数字日志-测试用例1",
			args: args{
				log: "dig1 8 1 5 1",
			},
			want: true,
		},
		{
			name: "是否是数字日志-测试用例2",
			args: args{
				log: "let1 art can",
			},
			want: false,
		},
		{
			name: "是否是数字日志-测试用例3",
			args: args{
				log: "let3 art zero",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isDigit(tt.args.log); got != tt.want {
				t.Errorf("isDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reorderLogFiles(t *testing.T) {
	type args struct {
		logs []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "重新排列日志文件-测试用例1",
			args: args{
				logs: []string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"},
			},
			want: []string{"let1 art can", "let3 art zero", "let2 own kit dig", "dig1 8 1 5 1", "dig2 3 6"},
		},
		{
			name: "重新排列日志文件-测试用例2",
			args: args{
				logs: []string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo"},
			},
			want: []string{"g1 act car", "a8 act zoo", "ab1 off key dog", "a1 9 2 3 1", "zo4 4 7"},
		},
		{
			name: "重新排列日志文件-测试用例2",
			args: args{
				logs: []string{"zoey i love you", "lucas i love you", "rong i love you"},
			},
			want: []string{"lucas i love you", "rong i love you", "zoey i love you"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reorderLogFiles(tt.args.logs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reorderLogFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}
