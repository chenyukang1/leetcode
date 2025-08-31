package 链表

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	slow, fast := head, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	if slow.Next == nil {
		return nil
	} else {
		slow.Next = slow.Next.Next
	}
	return head
}
