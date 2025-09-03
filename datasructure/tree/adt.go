package tree

type Comparable[T any] interface {
	CompareTo(other T) int
}

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
