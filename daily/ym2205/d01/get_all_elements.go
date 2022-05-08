package d01

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) (ans []int) {

	valSli1 := make([]int, 0)
	valSli2 := make([]int, 0)

	var midOrder func(root *TreeNode, sli *[]int)
	midOrder = func(root *TreeNode, sli *[]int) {
		if root != nil {
			midOrder(root.Left, sli)
			*sli = append(*sli, root.Val)
			midOrder(root.Right, sli)
		}
	}
	midOrder(root1, &valSli1)
	midOrder(root2, &valSli2)

	if len(valSli1) == 0 {
		return valSli2
	}
	if len(valSli2) == 0 {
		return valSli1
	}

	i, j := 0, 0
outer:
	for i < len(valSli1) {
	inter:
		for j <= len(valSli2) {
			if valSli1[i] >= valSli2[j] {
				ans = append(ans, valSli2[j])
				j++
				continue outer
			}
			ans = append(ans, valSli1[i])
			i++
			continue inter
		}
	}
	return
}
