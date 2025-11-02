package 链表

// https://leetcode.cn/problems/swap-nodes-in-pairs/?envType=study-plan-v2&envId=top-100-liked

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head}
	prev, cur := dummyHead, dummyHead.Next
	for i := 0; cur != nil && cur.Next != nil; i++ {
		if i%2 == 0 {
			next := cur.Next
			prev.Next = next
			temp := next.Next
			next.Next = cur
			cur.Next = temp
			prev = next
		} else {
			prev = cur
			cur = cur.Next
		}
	}
	return dummyHead.Next
}
