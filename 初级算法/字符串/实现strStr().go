package 字符串

// 超时
func strStr(haystack string, needle string) int {
	for i, j := 0, 0; i < len(haystack); i++ {
		for j < len(needle) && i < len(haystack) && haystack[i] == needle[j] {
			i++
			j++
		}
		if j == len(needle) {
			return i - len(needle) - 1
		} else {
			j = 0
		}
	}
	return -1
}
