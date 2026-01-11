package 二叉树

// https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/?envType=study-plan-v2&envId=top-100-liked

func flatten(root *TreeNode) {
	var arr []*TreeNode
	preDsf(root, &arr)
	for i := 0; i < len(arr); i++ {
		arr[i].Left = nil
		if i == len(arr)-1 {
			arr[i].Right = nil
		} else {
			arr[i].Right = arr[i+1]
		}
	}
}

func preDsf(root *TreeNode, arr *[]*TreeNode) {
	if root == nil {
		return
	}
	*arr = append(*arr, root)
	preDsf(root.Left, arr)
	preDsf(root.Right, arr)
}
