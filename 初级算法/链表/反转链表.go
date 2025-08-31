package 链表

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	prev, cur := head, head.Next
	for cur != nil {
		temp := cur.Next
		cur.Next = prev
		prev = cur
		cur = temp
	}
	return prev
}
