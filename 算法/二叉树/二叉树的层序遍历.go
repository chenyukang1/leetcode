package 二叉树

func levelOrder(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		var arr []int
		for i := len(queue); i > 0; i-- {
			node := queue[0]
			queue = queue[1:]
			arr = append(arr, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, arr)
	}
	return
}
