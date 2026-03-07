# https://leetcode.cn/problems/target-sum/description/

from typing import List


class Solution:
    def findTargetSumWays(self, nums: List[int], target: int) -> int:
        # left - right = target
        # left + right = sum(nums)
        # left = (target + sum(nums)) / 2
        tmp = target + sum(nums)
        if tmp % 2 == 1:
            return 0
        left = int(tmp / 2)

        # dp[i] 表示和为i的组合数
        dp = [0] * (left + 1)
        dp[0] = nums.count(0)
        for i in range(1, len(nums)):
            for j in range(target, nums[i] - 1, -1):
                dp[j] = max(dp[j], dp[j - nums[i]] + 1)

        return dp[left]
