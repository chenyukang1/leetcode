package 链表

func isPalindrome(head *ListNode) bool {
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	for i, j := 0, len(arr)-1; i < j; {
		if arr[i] != arr[j] {
			return false
		}
		i++
		j--
	}
	return true
}
