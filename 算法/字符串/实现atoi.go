package 字符串

func myAtoi(s string) int {
	i := 0
	sign := 1

	for i < len(s) && s[i] == ' ' {
		i++
	}
	if i == len(s) {
		return 0
	}

	if s[i] == '-' {
		sign = -1
	}

	for i < len(s) && s[i] == '0' {
		i++
	}
	if i == len(s) {
		return 0
	}

	sum := 0
	m := 1
	for j := len(s) - 1; j >= i; j-- {
		sum += m * (int)(s[j]-'0')
		m *= 10
	}
	return sign * sum
}
