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

    def rob2(self, nums: List[int]) -> int:
        # dp[i]表示到i房屋为止，能偷到的最大金额
        # dp[i] = max(dp[i-2]+nums[i],dp[i-1])
        if len(nums) == 1:
            return nums[0]
        dp = [0] * (len(nums))
        dp[0] = nums[0]
        dp[1] = max(nums[0], nums[1])

        for i in range(2, len(nums)):
            dp[i] = max(dp[i - 2] + nums[i], dp[i - 1])

        return dp[len(nums) - 1]
