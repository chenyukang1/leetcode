package 二叉树

// https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/?envType=study-plan-v2&envId=top-100-liked

func sortedArrayToBST(nums []int) *TreeNode {
	return recurseBST(0, len(nums)-1, nums)
}

func recurseBST(left int, right int, nums []int) *TreeNode {
	if left > right {
		return nil
	}
	mid := left + (right-left)>>1
	root := &TreeNode{nums[mid], nil, nil}
	root.Left = recurseBST(left, mid-1, nums)
	root.Right = recurseBST(mid+1, right, nums)
	return root
}
