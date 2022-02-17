package wk280

import "testing"

func Test_minimumRemoval(t *testing.T) {
	type args struct {
		beans []int
	}
	tests := []struct {
		name      string
		args      args
		wantCount int64
	}{
		// TODO: Add test cases.
		{
			name: "testcase1",
			args: args{
				beans: []int{4, 1, 6, 5},
			},
			wantCount: 4,
		},
		{
			name: "testcase2",
			args: args{
				beans: []int{2, 10, 3, 2},
			},
			wantCount: 7,
		},
		{
			name:      "testcase3",
			args:      args{beans: []int{63, 43, 12, 94}},
			wantCount: 83,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := minimumRemoval(tt.args.beans); gotCount != tt.wantCount {
				t.Errorf("minimumRemoval() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
