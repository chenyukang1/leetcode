package 树

import "math"

func isValidBST(root *TreeNode) bool {
	return isValidBST0(root, math.MinInt64, math.MaxInt64)
}

func isValidBST0(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}
	if root.Val >= max || root.Val <= min {
		return false
	}
	if root.Left != nil && root.Left.Val >= root.Val {
		return false
	}
	if root.Right != nil && root.Right.Val <= root.Val {
		return false
	}
	return isValidBST0(root.Left, min, root.Val) && isValidBST0(root.Right, root.Val, max)
}
