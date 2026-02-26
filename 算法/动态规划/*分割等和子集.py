from functools import cache
from typing import List


class Solution:
    def canPartition(self, nums: List[int]) -> bool:
        all_sum = sum(nums)
        if all_sum % 2 == 1:
            return False
        target = all_sum / 2

        nums.sort(reverse=True)

        @cache
        def back_track(start, sum):
            if sum == target:
                return True
            for i in range(start, len(nums)):
                if sum + nums[i] > target:
                    continue
                if back_track(i + 1, sum + nums[i]):
                    return True
            return False

        return back_track(0, 0)

    def canPartition2(self, nums: List[int]) -> bool:
        all_sum = sum(nums)
        if all_sum % 2 == 1:
            return False
        target = all_sum // 2
        n = len(nums)
        dp = [0] * (all_sum + 1)
        for i in range(n):
            for j in range(target, nums[i], -1):
                dp[j] = max(dp[j], dp[j - nums[i]] + nums[i])

        return dp[target] == target


if __name__ == "__main__":
    s = Solution()
    print(s.canPartition([2, 2, 3, 5]))
