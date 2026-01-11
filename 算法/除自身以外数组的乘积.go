package hot100

// https://leetcode.cn/problems/product-of-array-except-self/?envType=study-plan-v2&envId=top-100-liked

// productExceptSelf 时间复杂度O(n), 空间复杂度O(n)
func productExceptSelf(nums []int) (ans []int) {
	m := len(nums)
	l, r := make([]int, m), make([]int, m)
	l[0] = 1
	r[m-1] = 1

	for i := 1; i < m; i++ {
		l[i] = l[i-1] * nums[i-1]
	}
	for i := m - 2; i >= 0; i-- {
		r[i] = r[i+1] * nums[i+1]
	}
	for i := 0; i < m; i++ {
		ans = append(ans, l[i]*r[i])
	}
	return
}

// productExceptSelf2 时间复杂度O(n), 空间复杂度O(1)
func productExceptSelf2(nums []int) []int {
	m := len(nums)
	ans := make([]int, m)
	ans[0] = 1

	for i := 1; i < m; i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}
	r := 1
	for i := m - 1; i >= 0; i-- {
		ans[i] *= r
		r *= nums[i]
	}
	return ans
}
