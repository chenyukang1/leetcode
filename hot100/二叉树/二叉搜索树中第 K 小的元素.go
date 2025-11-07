package 二叉树

// https://leetcode.cn/problems/kth-smallest-element-in-a-bst/?envType=study-plan-v2&envId=top-100-liked

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	var arr []int
	dsf(root, &arr)
	return arr[k-1]
}

func dsf(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	dsf(root.Left, arr)
	*arr = append(*arr, root.Val)
	dsf(root.Right, arr)
}
