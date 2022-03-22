package d19

import "testing"

var (
	root = TreeNode{
		val: 1,
		left: &TreeNode{
			val: 2,
			left: &TreeNode{
				val:   4,
				left:  nil,
				right: nil,
			},
			right: nil,
		},
		right: &TreeNode{
			val:   3,
			left:  nil,
			right: nil,
		},
	}
)

func Test_tree2str(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantRet string
	}{
		{
			name: "根据二叉树创建字符串-测试用例1",
			args: args{
				root: &root,
			},
			wantRet: "1(2(4))(3)",
		},
		{
			name: "根据二叉树创建字符串-测试用例2",
			args: args{
				root: &root,
			},
			wantRet: "1(2(4))(3)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRet := tree2str(tt.args.root); gotRet != tt.wantRet {
				t.Errorf("tree2str() = %v, want %v", gotRet, tt.wantRet)
			}
		})
	}
}

func Test_tree2str2(t *testing.T) {

}
