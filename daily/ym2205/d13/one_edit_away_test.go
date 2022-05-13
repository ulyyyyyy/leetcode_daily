package d13

import (
	"math/rand"
	"strconv"
	"testing"
)

func Test_oneEditAway(t *testing.T) {
	type args struct {
		first  string
		second string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "面试题 01.05. 一次编辑-测试用例1",
			args: args{
				first:  "pale",
				second: "ple",
			},
			want: true,
		},
		{
			name: "面试题 01.05. 一次编辑-测试用例2",
			args: args{
				first:  "pales",
				second: "pal",
			},
			want: false,
		},
		{
			name: "面试题 01.05. 一次编辑-测试用例3",
			args: args{
				first:  "abcdefg",
				second: "abcdef",
			},
			want: true,
		},
		{
			name: "面试题 01.05. 一次编辑-测试用例4",
			args: args{
				first:  "",
				second: "",
			},
			want: true,
		},
		{
			name: "面试题 01.05. 一次编辑-测试用例4",
			args: args{
				first:  "",
				second: "a",
			},
			want: true,
		},
		{
			name: "面试题 01.05. 一次编辑-测试用例5",
			args: args{
				first:  "teacher",
				second: "bleacher",
			},
			want: false,
		},
		{
			name: "面试题 01.05. 一次编辑-测试用例5",
			args: args{
				first:  "islander",
				second: "slander",
			},
			want: true,
		},
		{
			name: "面试题 01.05. 一次编辑-测试用例6",
			args: args{
				first:  "islander",
				second: "slander",
			},
			want: true,
		},
		{
			name: "面试题 01.05. 一次编辑-测试用例6",
			args: args{
				first:  "islande",
				second: "islander",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oneEditAway(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("oneEditAway() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkOneEditAway(b *testing.B) {
	for i := 0; i < b.N; i++ {
		oneEditAway("pales", "ple")
		oneEditAway("abcdefg", "abcdefg")
		oneEditAway("horse", "ros")
		oneEditAway("", "a")
		oneEditAway("teacher", "bleacher")
		oneEditAway("islander", "slander")
	}
}

func BenchmarkOneEditAway2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		oneEditAway2("pales", "ple")
		oneEditAway2("abcdefg", "abcdefg")
		oneEditAway2("horse", "ros")
		oneEditAway2("", "a")
		oneEditAway2("teacher", "bleacher")
		oneEditAway2("islander", "slander")
	}
}

func Benchmark_CompareStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := strconv.Itoa(rand.Int())
		compareStr(s, s)
	}
}

func Benchmark_CompareStrByBit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := strconv.Itoa(rand.Int())
		compareStrByBit(s, s)
	}
}
