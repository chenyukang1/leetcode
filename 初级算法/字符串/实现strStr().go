package 字符串

// 暴力搜索，超时
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

// KMP算法 O(n * m)
func strStr2(haystack string, needle string) int {
	next := getNext(needle)
	for i, j := 0, 0; i < len(haystack); i++ {
		for j < len(needle) && i < len(haystack) && haystack[i] == needle[j] {
			i++
			j++
		}
		if j == len(needle) {
			return i - len(needle) - 1
		} else {
			j = next[j-1]
		}
	}
	return -1
}

func getNext(needle string) []int {
	next := make([]int, len(needle))
	j := -1
	next[0] = -1
	for i := 1; i < len(needle); i++ {
		for j > 0 && needle[j+1] != needle[i] {
			j = next[j]
		}
		if needle[j+1] == needle[i] {
			j++
		}
		next[i] = j
	}
	return next
}
