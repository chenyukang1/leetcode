package hot100

func findAnagrams(s string, p string) []int {
	var ans []int
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return ans
	}

	var sCount, pCount [26]int
	for i, ch := range p {
		sCount[s[i]-'a']++
		pCount[ch-'a']++
	}
	if sCount == pCount {
		ans = append(ans, 0)
	}

	for i := pLen; i < sLen; i++ {
		sCount[s[i-pLen]-'a']--
		sCount[s[i]-'a']++
		if sCount == pCount {
			ans = append(ans, i-pLen+1)
		}
	}
	return ans
}
