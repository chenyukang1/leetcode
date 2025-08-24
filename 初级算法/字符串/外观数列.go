package 字符串

import (
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	// dp[0] = 1
	// dp[i] 表示外观数列第i+1个元素
	// dp[i] = countAndSay(dp[i-1])
	dp := make([]string, n)
	dp[0] = "1"
	for i := 1; i < n; i++ {
		str := dp[i-1]
		count := 1
		var builder strings.Builder
		for j := 0; j < len(str); j++ {
			for ; j < len(str)-1 && str[j] == str[j+1]; j++ {
				count++
			}
			builder.WriteString(strconv.Itoa(count))
			builder.WriteByte(str[j])
			count = 1
		}
		dp[i] = builder.String()
	}
	return dp[n-1]
}

func countAndSay2(n int) string {
	if n == 1 {
		return "1"
	}
	str := countAndSay2(n - 1)
	count := 1
	var builder strings.Builder
	for i := 0; i < len(str); i++ {
		for ; i < len(str)-1 && str[i] == str[i+1]; i++ {
			count++
		}
		builder.WriteString(strconv.Itoa(count))
		builder.WriteByte(str[i])
		count = 1
	}
	return builder.String()
}
