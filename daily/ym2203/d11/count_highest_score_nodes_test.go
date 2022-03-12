package d11

import "testing"

func Test_countHighestScoreNodes(t *testing.T) {
	type args struct {
		parents []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "",
			args: args{
				parents: []int{-1, 2, 0, 2, 0},
			},
			wantAns: 3,
		},
		{
			name: "",
			args: args{
				parents: []int{-1, 2, 0},
			},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := countHighestScoreNodesOffice(tt.args.parents); gotAns != tt.wantAns {
				t.Errorf("countHighestScoreNodesOffice() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
