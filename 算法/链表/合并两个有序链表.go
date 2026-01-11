package 链表

// https://leetcode.cn/problems/merge-two-sorted-lists/?envType=study-plan-v2&envId=top-100-liked

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	p := head
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			p.Next = list1
			p = p.Next
			list1 = list1.Next
		} else {
			p.Next = list2
			p = p.Next
			list2 = list2.Next
		}
	}
	if list1 != nil {
		p.Next = list1
	} else {
		p.Next = list2
	}
	return head.Next
}
