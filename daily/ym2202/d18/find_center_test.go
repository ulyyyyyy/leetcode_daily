package d18

import "testing"

func Test_findCenter(t *testing.T) {
	type args struct {
		edges [][]int
	}
	tests := []struct {
		name     string
		args     args
		wantNode int
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNode := findCenter(tt.args.edges); gotNode != tt.wantNode {
				t.Errorf("findCenter() = %v, want %v", gotNode, tt.wantNode)
			}
		})
	}
}
