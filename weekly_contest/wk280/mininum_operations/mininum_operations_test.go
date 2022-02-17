package mininum_operations

import "testing"

func Test_minimumOperations(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
	}{
		{
			name: "testcase1",
			args: args{
				nums: []int{3, 1, 3, 2, 4, 3},
			},
			wantCount: 3,
		},
		{
			name: "testcase2",
			args: args{
				nums: []int{1, 2, 2, 2, 2},
			},
			wantCount: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := minimumOperations(tt.args.nums); gotCount != tt.wantCount {
				t.Errorf("minimumOperations() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
