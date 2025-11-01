package 链表

// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/?envType=study-plan-v2&envId=top-100-liked

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l := 0
	p := head
	for p != nil {
		l++
		p = p.Next
	}
	dummyHead := &ListNode{Val: 0, Next: head}
	cur := dummyHead
	for i := 0; i < l-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummyHead.Next
}
