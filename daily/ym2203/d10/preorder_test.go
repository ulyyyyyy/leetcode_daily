package d10

import (
	"reflect"
	"testing"
)

var root = &Node{
	Val: 1,
	Children: []*Node{
		{Val: 3, Children: []*Node{
			{Val: 5, Children: make([]*Node, 0)}, {Val: 6, Children: make([]*Node, 0)}},
		},
		{Val: 2, Children: make([]*Node, 0)},
		{Val: 4, Children: make([]*Node, 0)},
	},
}

func Test_preOrder(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name       string
		args       args
		wantValSli []int
	}{
		{
			name: "N 叉树的前序遍历-测试用例1",
			args: args{
				root: root,
			},
			wantValSli: []int{1, 3, 5, 6, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValSli := preOrder(tt.args.root); !reflect.DeepEqual(gotValSli, tt.wantValSli) {
				t.Errorf("preOrder() = %v, want %v", gotValSli, tt.wantValSli)
			}
		})
	}
}

func Test_postOrder(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name       string
		args       args
		wantValSli []int
	}{
		{
			name: "N 叉树的后序遍历-测试用例1",
			args: args{
				root: root,
			},
			wantValSli: []int{5, 6, 3, 2, 4, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValSli := postOrder(tt.args.root); !reflect.DeepEqual(gotValSli, tt.wantValSli) {
				t.Errorf("postOrder() = %v, want %v", gotValSli, tt.wantValSli)
			}
		})
	}
}

func Test_preOrder1(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name       string
		args       args
		wantValSli []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValSli := preOrder(tt.args.root); !reflect.DeepEqual(gotValSli, tt.wantValSli) {
				t.Errorf("preOrder() = %v, want %v", gotValSli, tt.wantValSli)
			}
		})
	}
}
