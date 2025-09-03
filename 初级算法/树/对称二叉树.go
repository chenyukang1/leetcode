package 树

// 层序遍历
func isSymmetric(root *TreeNode) bool {
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		var tmp []*TreeNode
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

// 递归
func isSymmetric2(root *TreeNode) bool {
	return isSymmetric0(root.Left, root.Right)
}

func isSymmetric0(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return isSymmetric0(left.Right, right.Left) && isSymmetric0(left.Left, right.Right)
}
