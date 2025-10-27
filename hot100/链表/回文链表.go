package 链表

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}
	var arr []*ListNode
	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}
	for l, r := 0, len(arr)-1; l < r; l, r = l+1, r-1 {
		if arr[l].Val != arr[r].Val {
			return false
		}
	}
	return true
}
