package d16

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor(root, p *TreeNode) *TreeNode {
	var st []*TreeNode
	var pre, cur *TreeNode = nil, root
	for len(st) > 0 || cur != nil {
		for cur != nil {
			st = append(st, cur)
			cur = cur.Left
		}
		cur = st[len(st)-1]
		st = st[:len(st)-1]
		if pre == p {
			return cur
		}
		pre = cur
		cur = cur.Right
	}
	return nil
}
