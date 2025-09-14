package 动态规划

import "math"

// 贪心：找区间内的最大差值
func maxProfit(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}
	m := math.MaxInt32
	ans := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < m {
			m = prices[i]
		}
		temp := prices[i] - m
		if temp > ans {
			ans = temp
		}
	}
	return ans
}

// 动态规划
func maxProfit2(prices []int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}
	// dp[i][0] 第i+1天持有股票的最多钱
	// dp[i][1] 第i+1天不持有股票的最多钱
	// dp[i][0] = max{dp[i-1][0], -prices[i]}
	// dp[i][1] = max{dp[i-1][1], dp[i-1][0] + prices[i]}
	dp := make([][2]int, len(prices))
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < len(prices); i++ {
		if dp[i-1][0] < -prices[i] {
			dp[i][0] = -prices[i]
		} else {
			dp[i][0] = dp[i-1][0]
		}

		temp := dp[i-1][0] + prices[i]
		if dp[i-1][1] > temp {
			dp[i][1] = dp[i-1][1]
		} else {
			dp[i][1] = temp
		}
	}

	return dp[len(prices)-1][1]
}
