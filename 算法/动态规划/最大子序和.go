package 动态规划

func maxSubArray(nums []int) int {
	// dp[i] 截止到i的最大和
	// dp[i] = max{dp[i-1]+nums[i], nums[i]}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
	}
	return dp[len(nums)-1]
}
