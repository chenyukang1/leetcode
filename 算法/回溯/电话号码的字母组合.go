package 回溯算法

//https://leetcode.cn/problems/letter-combinations-of-a-phone-number/description/

func letterCombinations(digits string) []string {
	m := map[byte][]byte{
		2: {'a', 'b', 'c'},
		3: {'d', 'e', 'f'},
		4: {'g', 'h', 'i'},
		5: {'j', 'k', 'l'},
		6: {'m', 'n', 'o'},
		7: {'p', 'q', 'r', 's'},
		8: {'t', 'u', 'v'},
		9: {'w', 'x', 'y', 'z'},
	}
	var ans []string
	var path []byte

	backTracking4(digits, m, 0, path, &ans)
	return ans
}

func backTracking4(digits string, m map[byte][]byte, start int, path []byte, ans *[]string) {
	if len(path) == len(digits) {
		*ans = append(*ans, string(path))
		return
	}

	for i := start; i < len(digits); i++ {
		digit, ok := m[digits[i]-'0']
		if ok {
			for _, s := range digit {
				path = append(path, s)
				backTracking4(digits, m, i+1, path, ans)
				path = path[:len(path)-1]
			}
		}
	}
}
