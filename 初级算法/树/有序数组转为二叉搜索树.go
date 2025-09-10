package 树

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return insert(nums, 0, len(nums))
}

// 把当前中间节点作为根节点，递归地创建左子树和右子树
// [left, right)
func insert(nums []int, left int, right int) *TreeNode {
	if left >= right {
		return nil
	}
	mid := left + ((right - left) >> 1)
	root := &TreeNode{Val: nums[mid]}
	root.Left = insert(nums, left, mid)
	root.Right = insert(nums, mid+1, right)
	return root
}
