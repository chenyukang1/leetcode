package hot100

func lengthOfLongestSubstring(s string) int {
	i, j, ans := 0, 0, 0
	m := make(map[byte]int)
	for j < len(s) {
		for _, ok := m[s[j]]; j < len(s) && !ok; j++ {
			m[s[j]] = j
		}
		ans = maxInt(ans, j-i+1)
		if j == len(s) {
			break
		}
		for i < m[s[j]] {
			delete(m, s[i])
			i++
		}
	}
	return ans
}
