package 链表

// https://leetcode.cn/problems/intersection-of-two-linked-lists/?envType=study-plan-v2&envId=top-100-liked

type ListNode struct {
	Val  int
	Next *ListNode
}

// getIntersectionNode 时间复杂度O(n)，空间复杂度O(n)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]struct{})
	for p := headA; p != nil; {
		m[p] = struct{}{}
		p = p.Next
	}
	for p := headB; p != nil; {
		if _, ok := m[p]; ok {
			return p
		}
		p = p.Next
	}
	return nil
}

// getIntersectionNode2 时间复杂度O(n)，空间复杂度O(1)
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pA, pB := headA, headB
	for pA != pB {
		if pA != nil {
			pA = pA.Next
		} else {
			pA = headB
		}
		if pB != nil {
			pB = pB.Next
		} else {
			pB = headA
		}
	}
	return pA
}
