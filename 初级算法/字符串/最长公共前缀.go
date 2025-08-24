package 字符串

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) < 2 {
		return strs[0]
	}
	var builder strings.Builder
	for i := 0; i < len(strs[0]); i++ {
		ans := strs[0][i]
		j := 1
		for ; j < len(strs); j++ {
			if len(strs[j])-1 < i {
				break
			}
			if ans != strs[j][i] {
				break
			}
		}
		if j != len(strs) {
			break
		}
		builder.WriteByte(strs[0][i])
	}
	return builder.String()
}
