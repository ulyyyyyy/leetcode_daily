package d12

// 给定一个 n 叉树的根节点 root ，返回 其节点值的 后序遍历 。
//
// n 叉树 在输入中按层序遍历进行序列化表示，每组子节点由空值 null 分隔（请参见示例）。

type Node struct {
	Val      int
	Children []*Node
}

func postOrder(root *Node) (valSli []int) {
	var post func(node *Node)

	post = func(node *Node) {
		if node == nil {
			return
		}
		if node.Children == nil {
			return
		}
		for _, child := range node.Children {
			post(child)
		}
		valSli = append(valSli, node.Val)
	}
	post(root)
	return
}
