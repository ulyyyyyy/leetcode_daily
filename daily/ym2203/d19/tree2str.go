package d19

import (
	"fmt"
	"strconv"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 你需要采用前序遍历(根/左/右)的方式，将一个二叉树转换成一个由括号和整数组成的字符串。
//
// 空节点则用一对空括号 "()" 表示。而且你需要省略所有不影响字符串与原始二叉树之间的一对一映射关系的空括号对。

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func tree2str(root *TreeNode) (ret string) {
	switch {
	case root == nil:
		return ""
	// 如果 root 左右子节点都不存在，则返回 root
	case root.left == nil && root.right == nil:
		return strconv.Itoa(root.val)
	// 如果 root 只有左节点存在，则返回 root(left)
	case root.right == nil:
		return fmt.Sprintf("%d(%s)", root.val, tree2str(root.left))
	// 如果 root 只有右节点 或 左、右节点都 存在，则返回 root(left)(right)。
	// 因为 left 可能为空，所以可能是 root()(right)
	default:
		return fmt.Sprintf("%d(%s)(%s)", root.val, tree2str(root.left), tree2str(root.right))
	}
}
