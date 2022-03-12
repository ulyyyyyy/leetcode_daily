package p17

import (
	"reflect"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name       string
		args       args
		wantRetSli []string
	}{
		{
			name: "电话号码的字母组合-测试用例1",
			args: args{
				digits: "23",
			},
			wantRetSli: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name: "电话号码的字母组合-测试用例2",
			args: args{
				digits: "",
			},
			wantRetSli: []string{},
		},
		{
			name: "电话号码的字母组合-测试用例3",
			args: args{
				digits: "2",
			},
			wantRetSli: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetSli := letterCombinations(tt.args.digits); !reflect.DeepEqual(gotRetSli, tt.wantRetSli) {
				t.Errorf("letterCombinations() = %v, want %v", gotRetSli, tt.wantRetSli)
			}
		})
	}
}

func Test_letterCombinationsQueue(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "电话号码的字母组合-测试用例1",
			args: args{
				digits: "23",
			},
			want: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name: "电话号码的字母组合-测试用例2",
			args: args{
				digits: "",
			},
			want: []string{},
		},
		{
			name: "电话号码的字母组合-测试用例3",
			args: args{
				digits: "2",
			},
			want: []string{"a", "b", "c"},
		},
		{
			name: "电话号码的字母组合-测试用例4",
			args: args{
				digits: "234",
			},
			want: []string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi", "bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi", "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := letterCombinationsQueue(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterCombinationsQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}
