package 树

func levelOrder(root *TreeNode) [][]int {
	queue := []*TreeNode{root}
	var ans [][]int
	for len(queue) > 0 {
		size := len(queue)
		tmp := make([]int, 0, size)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[:1]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			tmp = append(tmp, node.Val)
		}
		ans = append(ans, tmp)
	}
	return ans
}
