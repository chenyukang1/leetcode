# https://leetcode.cn/problems/house-robber/?envType=study-plan-v2&envId=top-100-liked

from typing import List


class Solution:
    def rob(self, nums: List[int]) -> int:
        # dp[i][0] = max(dp[i-1][0], dp[i-1][1])
        # dp[i][1] = dp[i-1][0] + nums[i]
        dp = [[0, 0] for i in range(len(nums))]
        dp[0][0] = 0
        dp[0][1] = nums[0]
        for i in range(1, len(nums)):
            dp[i][0] = max(dp[i - 1][0], dp[i - 1][1])
            dp[i][1] = dp[i - 1][0] + nums[i]

        return max(dp[-1][0], dp[-1][1])
