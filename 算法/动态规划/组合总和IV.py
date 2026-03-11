# https://leetcode.cn/problems/combination-sum-iv/
from typing import List


class Solution:
    def combinationSum4(self, nums: List[int], target: int) -> int:
        # dp[i] 表示总和为i的排列数
        dp = [0] * (target + 1)
        dp[0] = 1
        # 先背包再物品
        for i in range(target + 1):
            for j in range(len(nums)):
                if i >= nums[j]:
                    dp[i] += dp[i - nums[j]]

        return dp[target]


if __name__ == "__main__":
    solution = Solution()
    print(solution.combinationSum4([1, 2, 3], 4))
