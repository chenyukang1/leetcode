package 二叉树

// https://leetcode.cn/problems/binary-tree-inorder-traversal/description/?envType=study-plan-v2&envId=top-100-liked

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var ans []int
	recurseTree(root, &ans)
	return ans
}

func recurseTree(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	recurseTree(root.Left, ans)
	*ans = append(*ans, root.Val)
	recurseTree(root.Right, ans)
}
