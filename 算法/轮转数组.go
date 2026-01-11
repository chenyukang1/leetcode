package hot100

// https://leetcode.cn/problems/rotate-array/?envType=study-plan-v2&envId=top-100-liked

func rotate(nums []int, k int) {
	l := len(nums)
	k = k % l
	temp := make([]int, k)
	copy(temp, nums[l-k:])
	for i := l - 1; i >= k; i-- {
		nums[i] = nums[i-k]
	}
	copy(nums[:k], temp)
}
