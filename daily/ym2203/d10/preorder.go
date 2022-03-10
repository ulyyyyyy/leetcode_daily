package d10

// 给定一个 n 叉树的根节点  root ，返回 其节点值的 前序遍历、中序遍历、后序遍历。
//
// n 叉树 在输入中按层序遍历进行序列化表示，每组子节点由空值 null 分隔。

type Node struct {
	Val      int
	Children []*Node
}

func preOrder(root *Node) (valSli []int) {
	var pre func(node *Node)

	pre = func(node *Node) {
		if node == nil {
			return
		}
		valSli = append(valSli, node.Val)
		if node.Children == nil {
			return
		}
		for _, child := range node.Children {
			pre(child)
		}
	}
	pre(root)
	return
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
