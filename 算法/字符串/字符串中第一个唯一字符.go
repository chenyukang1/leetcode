package 字符串

func firstUniqChar(s string) int {
	m := make(map[int32]int)
	for _, r := range s {
		m[r]++
	}
	for i, r := range s {
		if m[r] == 1 {
			return i
		}
	}

	return -1
}
