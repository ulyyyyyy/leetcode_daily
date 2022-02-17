package count_operations

import "testing"

func Test_countOperations(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
	}{
		{
			name: "testcase-1",
			args: args{
				num1: 2,
				num2: 3,
			},
			wantCount: 3,
		},
		{
			name: "testcase-2",
			args: args{
				num1: 10,
				num2: 10,
			},
			wantCount: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := countOperations(tt.args.num1, tt.args.num2); gotCount != tt.wantCount {
				t.Errorf("countOperations() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
