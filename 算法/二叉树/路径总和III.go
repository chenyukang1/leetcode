package 二叉树

// https://leetcode.cn/problems/path-sum-iii/?envType=study-plan-v2&envId=top-100-liked

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	return countPath(root, targetSum) + pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum)
}

func countPath(root *TreeNode, target int) int {
	if root == nil {
		return 0
	}

	count := 0
	if root.Val == target {
		count++
	}
	return count + countPath(root.Left, target-root.Val) + countPath(root.Right, target-root.Val)
}
