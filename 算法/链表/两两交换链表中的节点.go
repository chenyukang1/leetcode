package 链表

// https://leetcode.cn/problems/swap-nodes-in-pairs/?envType=study-plan-v2&envId=top-100-liked

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummyHead := &ListNode{0, head}
	prev, cur, next := dummyHead, dummyHead.Next, dummyHead.Next.Next
	for i := 0; next != nil; i++ {
		if i%2 == 0 {
			prev.Next = next
			temp := next.Next
			next.Next = cur
			cur.Next = temp
			prev = next
			next = temp
		} else {
			prev = cur
			cur = next
			next = next.Next
		}
	}
	return dummyHead.Next
}
