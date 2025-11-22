package 二叉树

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	inorderMap := make(map[int]int)
	for i, val := range inorder {
		inorderMap[val] = i
	}
	return buildTreeRecurse(0, len(preorder)-1, 0, len(inorder)-1, preorder, inorderMap)
}

func buildTreeRecurse(pStart int, pEnd int, iStart int, iEnd int, preorder []int, inorderMap map[int]int) *TreeNode {
	if pStart > pEnd {
		return nil
	}
	root := &TreeNode{Val: preorder[pStart]}
	leftSize := inorderMap[preorder[pStart]] - iStart
	root.Left = buildTreeRecurse(pStart+1, pStart+leftSize, iStart, inorderMap[preorder[pStart]]-1, preorder, inorderMap)
	root.Right = buildTreeRecurse(pStart+leftSize+1, pEnd, inorderMap[preorder[pStart]]+1, iEnd, preorder, inorderMap)
	return root
}
