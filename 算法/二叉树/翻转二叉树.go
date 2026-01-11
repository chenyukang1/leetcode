package 二叉树

// https://leetcode.cn/problems/invert-binary-tree/?envType=study-plan-v2&envId=top-100-liked

func invertTree(root *TreeNode) *TreeNode {
	recurseInvert(root)
	return root
}

func recurseInvert(root *TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	recurseInvert(root.Left)
	recurseInvert(root.Right)
}
