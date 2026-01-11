package 二叉树

// https://leetcode.cn/problems/maximum-depth-of-binary-tree/?envType=study-plan-v2&envId=top-100-liked

func maxDepth(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}
	return maxInt(dfs(root.Left, depth+1), dfs(root.Right, depth+1))
}

func maxInt(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
