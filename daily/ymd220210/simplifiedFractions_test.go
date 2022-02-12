package ymd220210

import (
	"reflect"
	"testing"
)

func Test_simplifiedFractions(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name     string
		args     args
		wantRlts []string
	}{
		{
			name:     "最简分数-测试用例1",
			args:     args{n: 2},
			wantRlts: []string{"1/2"},
		},
		{
			name:     "最简分数-测试用例2",
			args:     args{n: 3},
			wantRlts: []string{"1/2", "1/3", "2/3"},
		},
		{
			name:     "最简分数-测试用例3",
			args:     args{n: 4},
			wantRlts: []string{"1/2", "1/3", "2/3", "1/4", "3/4"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRlts := simplifiedFractions(tt.args.n); !reflect.DeepEqual(gotRlts, tt.wantRlts) {
				t.Errorf("simplifiedFractions() = %v, want %v", gotRlts, tt.wantRlts)
			}
		})
	}
}
