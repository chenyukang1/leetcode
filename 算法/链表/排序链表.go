package 链表

// https://leetcode.cn/problems/sort-list/?envType=study-plan-v2&envId=top-100-liked

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	next := slow.Next
	slow.Next = nil
	left := sortList(head)
	right := sortList(next)

	dummyHead := &ListNode{}
	p := dummyHead
	for left != nil && right != nil {
		if left.Val < right.Val {
			p.Next = left
			left = left.Next
		} else {
			p.Next = right
			right = right.Next
		}
		p = p.Next
	}
	if left == nil {
		p.Next = right
	}
	if right == nil {
		p.Next = left
	}
	return dummyHead.Next
}
