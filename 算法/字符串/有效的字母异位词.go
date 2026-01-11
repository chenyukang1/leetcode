package 字符串

func isAnagram(s string, t string) bool {
	arr := [26]int{}
	for _, i := range s {
		arr[i-'a']++
	}
	for _, i := range t {
		arr[i-'a']--
	}
	for _, i := range arr {
		if i != 0 {
			return false
		}
	}
	return true
}
