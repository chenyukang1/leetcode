package tree

import "testing"

type IntWrapper struct {
	Value int
}

func (a IntWrapper) CompareTo(b IntWrapper) int {
	if a.Value < b.Value {
		return -1
	} else if a.Value > b.Value {
		return 1
	}
	return 0
}

func TestContains(t *testing.T) {
	root := &BinaryNode[IntWrapper]{
		Data: IntWrapper{Value: 4},
	}
	l := &BinaryNode[IntWrapper]{
		Data: IntWrapper{Value: 2},
	}
	r := &BinaryNode[IntWrapper]{
		Data: IntWrapper{Value: 6},
	}
	root.Left = l
	root.Right = r
	if !Contains(IntWrapper{Value: 2}, root) {
		t.Errorf("expected true, got false!")
	}
	if Contains(IntWrapper{Value: 3}, root) {
		t.Errorf("expected false, got true!")
	}
}

func TestInsert(t *testing.T) {
	root := &BinaryNode[IntWrapper]{
		Data: IntWrapper{Value: 4},
	}
	l := &BinaryNode[IntWrapper]{
		Data: IntWrapper{Value: 2},
	}
	r := &BinaryNode[IntWrapper]{
		Data: IntWrapper{Value: 6},
	}
	root.Left = l
	root.Right = r
	Insert(IntWrapper{Value: 3}, root)
	value := root.Left.Right.Data.Value
	if value != 3 {
		t.Errorf("expected 3, got %d!", value)
	}
}

func TestRemove(t *testing.T) {
	root := &BinaryNode[IntWrapper]{
		Data: IntWrapper{Value: 6},
	}
	n1 := &BinaryNode[IntWrapper]{
		Data: IntWrapper{Value: 2},
	}
	n2 := &BinaryNode[IntWrapper]{
		Data: IntWrapper{Value: 8},
	}
	n3 := &BinaryNode[IntWrapper]{
		Data: IntWrapper{Value: 1},
	}
	n4 := &BinaryNode[IntWrapper]{
		Data: IntWrapper{Value: 3},
	}
	n5 := &BinaryNode[IntWrapper]{
		Data: IntWrapper{Value: 4},
	}
	root.Left = n1
	root.Right = n2
	n1.Left = n3
	n1.Right = n5
	n5.Left = n4
	Remove(IntWrapper{Value: 3}, root)
	Print(root)
	if n5.Left != nil {
		t.Errorf("left expected nil, got %d!", n5.Left.Data)
	}

	Remove(IntWrapper{Value: 4}, root)
	Print(root)
	if n1.Right != nil {
		t.Errorf("right expected nil, got %d!", n1.Right.Data)
	}
}

func TestRemove_2node(t *testing.T) {
	root := &BinaryNode[IntWrapper]{
		Data: IntWrapper{Value: 6},
	}
	n1 := &BinaryNode[IntWrapper]{
		Data: IntWrapper{Value: 2},
	}
	n2 := &BinaryNode[IntWrapper]{
		Data: IntWrapper{Value: 8},
	}
	n3 := &BinaryNode[IntWrapper]{
		Data: IntWrapper{Value: 1},
	}
	n4 := &BinaryNode[IntWrapper]{
		Data: IntWrapper{Value: 3},
	}
	n5 := &BinaryNode[IntWrapper]{
		Data: IntWrapper{Value: 4},
	}
	root.Left = n1
	root.Right = n2
	n1.Left = n3
	n1.Right = n5
	n5.Left = n4
	Remove(IntWrapper{Value: 2}, root)
	Print(root)
	if root.Left.Data.Value != 3 {
		t.Errorf("left expected nil, got %d!", root.Left.Data.Value)
	}
}
