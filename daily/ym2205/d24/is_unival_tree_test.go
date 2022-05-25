package d24

import (
	"sync"
	"testing"
)

var treePool1 = sync.Pool{
	New: func() any {
		return &TreeNode{Val: 1}
	},
}

var treePool2 = sync.Pool{
	New: func() any {
		return &TreeNode{Val: 2}
	},
}

func Test_isUnivalTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "测试用例1",
			args: args{
				root: &TreeNode{
					Val:   1,
					Left:  treePool1.Get().(*TreeNode),
					Right: treePool1.Get().(*TreeNode),
				},
			},
			want: true,
		},
		{
			name: "测试用例3",
			args: args{
				root: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   2,
						Left:  treePool2.Get().(*TreeNode),
						Right: treePool2.Get().(*TreeNode),
					},
					Right: &TreeNode{
						Val:   1,
						Left:  treePool2.Get().(*TreeNode),
						Right: treePool2.Get().(*TreeNode),
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUnivalTree(tt.args.root); got != tt.want {
				t.Errorf("isUnivalTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
