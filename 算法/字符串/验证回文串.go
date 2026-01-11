package 字符串

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	for l, r := 0, len(s)-1; l < r; {
		if s[l] < 'a' || s[l] > 'z' {
			l++
		}
		if s[r] < 'a' || s[r] > 'z' {
			r--
		}
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
