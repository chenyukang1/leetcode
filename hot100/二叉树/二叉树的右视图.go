package 二叉树

func rightSideView(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if i == size-1 {
				ans = append(ans, node.Val)
			}
		}
	}
	return ans
}
