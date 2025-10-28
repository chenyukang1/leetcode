package 链表

// https://leetcode.cn/problems/add-two-numbers/?envType=study-plan-v2&envId=top-100-liked

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	p := dummyHead
	add := 0
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + add
		add = sum / 10
		l1 = l1.Next
		l2 = l2.Next
		p.Next = &ListNode{Val: sum % 10}
		p = p.Next
	}
	for l1 != nil {
		sum := l1.Val + add
		add = sum / 10
		l1 = l1.Next
		p.Next = &ListNode{Val: sum % 10}
		p = p.Next
	}
	for l2 != nil {
		sum := l2.Val + add
		add = sum / 10
		l2 = l2.Next
		p.Next = &ListNode{Val: sum % 10}
		p = p.Next
	}
	if add != 0 {
		p.Next = &ListNode{Val: add}
		p = p.Next
	}
	return dummyHead.Next
}
