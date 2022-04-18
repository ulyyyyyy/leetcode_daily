package d18

import (
	"reflect"
	"testing"
)

func Test_lexicalOrder(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name      string
		args      args
		wantSlice []int
	}{
		{
			name: "字典序排数-测试用例1",
			args: args{
				n: 13,
			},
			wantSlice: []int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name: "字典序排数-测试用例2",
			args: args{
				n: 2,
			},
			wantSlice: []int{1, 2},
		},
		{
			name: "字典序排数-测试用例3",
			args: args{
				n: 50000,
			},
			wantSlice: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSlice := lexicalOrder(tt.args.n); !reflect.DeepEqual(gotSlice, tt.wantSlice) {
				t.Errorf("lexicalOrder() = %v, want %v", gotSlice, tt.wantSlice)
			}
		})
	}
}
