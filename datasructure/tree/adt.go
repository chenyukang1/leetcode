package tree

import "fmt"

type Comparable[T any] interface {
	CompareTo(other T) int
}

// BinaryNode 二叉查找树 ADT
type BinaryNode[T Comparable[T]] struct {
	Left  *BinaryNode[T]
	Right *BinaryNode[T]
	Data  T
}

func Contains[T Comparable[T]](data T, node *BinaryNode[T]) bool {
	if node == nil {
		return false
	}
	res := data.CompareTo(node.Data)
	if res == 0 {
		return true
	} else if res < 0 {
		return Contains(data, node.Left)
	} else {
		return Contains(data, node.Right)
	}
}

func Insert[T Comparable[T]](data T, node *BinaryNode[T]) *BinaryNode[T] {
	if node == nil {
		return &BinaryNode[T]{Data: data}
	}
	res := data.CompareTo(node.Data)
	if res < 0 {
		node.Left = Insert(data, node.Left)
	} else if res > 0 {
		node.Right = Insert(data, node.Right)
	} else {
		// duplicate
	}
	return node
}

func Remove[T Comparable[T]](data T, node *BinaryNode[T]) *BinaryNode[T] {
	if node == nil {
		return node // not found
	}
	res := data.CompareTo(node.Data)
	if res < 0 {
		node.Left = Remove(data, node.Left)
	} else if res > 0 {
		node.Right = Remove(data, node.Right)
	} else if node.Left != nil && node.Right != nil {
		// 有两个叶子节点的情况：先找到右子树的最小值，即比当前根节点大的最小值
		// 替换当前根节点的值，递归删除右子树的那个值
		node.Data = findMin(node.Right).Data
		node.Right = Remove(node.Data, node.Right)
	} else {
		if node.Left == nil {
			node = node.Right
		} else {
			node = node.Left
		}
	}
	return node
}

func findMin[T Comparable[T]](node *BinaryNode[T]) *BinaryNode[T] {
	if node == nil {
		return nil
	}
	if node.Left == nil {
		return node
	}
	return findMin(node.Left)
}

func Print[T Comparable[T]](root *BinaryNode[T]) {
	if root == nil {
		return
	}
	queue := []BinaryNode[T]{*root}
	var ans [][]T
	for len(queue) > 0 {
		size := len(queue)
		var temp []T
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, *node.Left)
			}
			if node.Right != nil {
				queue = append(queue, *node.Right)
			}
			temp = append(temp, node.Data)
		}
		ans = append(ans, temp)
	}
	fmt.Println(ans)
}
