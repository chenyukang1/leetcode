package 二叉树

// https://leetcode.cn/problems/symmetric-tree/description/?envType=study-plan-v2&envId=top-100-liked

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return isSymmetricRecurse(root.Left, root.Right)
}

func isSymmetricRecurse(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return isSymmetricRecurse(left.Left, right.Right) && isSymmetricRecurse(left.Right, right.Left)
}
