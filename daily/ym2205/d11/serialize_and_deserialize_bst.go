package d11

import (
	"math"
	"strconv"
	"strings"
)

// TreeNode is a Binary Search Tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Codec 这个类...不知道是干嘛的
type Codec struct {
}

func Constructor() (_ Codec) {
	return
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	ret := ""

	var postOrder func(root *TreeNode)
	postOrder = func(root *TreeNode) {
		if root != nil {
			postOrder(root.Left)
			postOrder(root.Right)
			ret += strconv.Itoa(root.Val)
		}
	}
	// 从根节点遍历
	postOrder(root)
	return ret
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	mid := strings.Split(data, " ")

	var construct func(lower int, upper int) *TreeNode
	construct = func(lower, upper int) *TreeNode {
		if len(mid) == 0 {
			return nil
		}
		val, _ := strconv.Atoi(mid[len(mid)-1])
		if val < lower || val > upper {
			return nil
		}
		mid = mid[:len(mid)-1]
		return &TreeNode{Val: val, Right: construct(val, upper), Left: construct(lower, val)}
	}
	return construct(math.MinInt32, math.MaxInt32)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
