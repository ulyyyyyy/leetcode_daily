package d21

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func findTarget(root *TreeNode, k int) (rlt bool) {
	tmpMap := make(map[int]struct{})
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if _, ok := tmpMap[k-root.Val]; ok {
			rlt = true
			return
		}
		// root.Val + n = k => n = k - root.Val
		tmpMap[root.Val] = struct{}{}
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return rlt
}
