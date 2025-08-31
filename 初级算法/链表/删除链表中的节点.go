package 链表

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	p, q := node, node
	for p.Next != nil {
		p.Val = p.Next.Val
		q = p
		p = p.Next
	}
	q.Next = nil
}
