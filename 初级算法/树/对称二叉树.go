package 树

func isSymmetric(root *TreeNode) bool {
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		tmp := make([]*TreeNode, size)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node != nil {
				queue = append(queue, node.Left)
				queue = append(queue, node.Right)
			}
			tmp = append(tmp, node)
		}
		for l, r := 0, size-1; l < r; {
			if tmp[l] == nil && tmp[r] == nil {
				l++
				r--
				continue
			}
			if tmp[l] == nil || tmp[r] == nil {
				return false
			}
			if tmp[l].Val != tmp[r].Val {
				return false
			}
			l++
			r--
		}
	}
	return true
}
