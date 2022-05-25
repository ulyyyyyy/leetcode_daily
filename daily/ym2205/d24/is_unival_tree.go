package d24

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	target := root.Val

	tmpSli := make([]*TreeNode, 1)
	tmpSli[0] = root

	for len(tmpSli) != 0 {
		s := len(tmpSli)
		for i := 0; i < s; i++ {
			if tmpSli[i].Val != target {
				return false
			}
			if tmpSli[i].Left != nil {
				tmpSli = append(tmpSli, tmpSli[i].Left)
			}
			if tmpSli[i].Right != nil {
				tmpSli = append(tmpSli, tmpSli[i].Right)
			}
		}
		tmpSli = tmpSli[s:]
	}
	return true
}
