package d09

import (
	"reflect"
	"testing"
)

func Test_diStringMatch(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "增减字符串匹配-测试用例1",
			args: args{
				s: "IDID",
			},
			wantAns: []int{0, 4, 1, 3, 2},
		},
		{
			name: "增减字符串匹配-测试用例1",
			args: args{
				s: "III",
			},
			wantAns: []int{0, 1, 2, 3},
		},
		{
			name: "增减字符串匹配-测试用例1",
			args: args{
				s: "DDI",
			},
			wantAns: []int{3, 2, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := diStringMatch(tt.args.s); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("diStringMatch() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
