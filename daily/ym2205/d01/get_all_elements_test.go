package d01

import (
	"reflect"
	"testing"
)

func Test_getAllElements(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{
			name: "两棵二叉搜索树中的所有元素",
			args: args{
				root1: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 4},
				},
				root2: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 0},
					Right: &TreeNode{Val: 3},
				},
			},
			wantAns: []int{0, 1, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := getAllElements(tt.args.root1, tt.args.root2); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("getAllElements() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
