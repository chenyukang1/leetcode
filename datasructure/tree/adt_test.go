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
