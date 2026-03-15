# https://leetcode.cn/problems/longest-increasing-subsequence/?envType=study-plan-v2&envId=top-100-liked

from typing import List


class Solution:
    def lengthOfLIS(self, nums: List[int]) -> int:
        ans = 1
        # dp[i]表示以nums[i]结尾的最长子序列
        # 为什么是nums[i]结尾？这样当nums[i]>nums[j]时，可以由 nums[j]+1递推
        dp = [1] * (len(nums))
        for i in range(1, len(nums)):
            for j in range(i):
                if nums[i] > nums[j]:
                    dp[i] = max(dp[i], dp[j] + 1)
            ans = max(ans, dp[i])

        return ans
