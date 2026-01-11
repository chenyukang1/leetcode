package 字符串

func reverseString(s []byte) {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		temp := s[l]
		s[l] = s[r]
		s[r] = temp
	}
}
