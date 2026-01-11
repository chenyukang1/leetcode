package 二叉树

// https://leetcode.cn/problems/diameter-of-binary-tree/?envType=study-plan-v2&envId=top-100-liked

func diameterOfBinaryTree(root *TreeNode) (ans int) {
	if root == nil {
		return 0
	}
	var depth func(root *TreeNode) int
	depth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := depth(root.Left) + 1
		right := depth(root.Right) + 1
		sum := left + right - 1
		if ans < sum {
			ans = sum
		}
		if left > right {
			return left
		}
		return right
	}
	depth(root)
	return ans - 1
}
