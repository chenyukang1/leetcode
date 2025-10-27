package 链表

var tail *ListNode

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	reverse(head, head.Next)
	return tail
}

func reverse(prev *ListNode, next *ListNode) {
	if next == nil {
		tail = prev
		return
	}
	reverse(next, next.Next)
	next.Next = prev
}
