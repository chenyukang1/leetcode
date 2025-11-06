package 二叉树

import "math"

// https://leetcode.cn/problems/validate-binary-search-tree/?envType=study-plan-v2&envId=top-100-liked

func isValidBST(root *TreeNode) bool {
	return isValidBSTRecurse(root, math.MaxInt64, math.MinInt64)
}

func isValidBSTRecurse(root *TreeNode, max int, min int) bool {
	if root == nil {
		return true
	}
	if root.Val >= max || root.Val <= min {
		return false
	}
	return isValidBSTRecurse(root.Left, root.Val, min) && isValidBSTRecurse(root.Right, max, root.Val)
}
