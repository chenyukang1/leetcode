# https://leetcode.cn/problems/maximum-product-subarray/?envType=study-plan-v2&envId=top-100-liked

from typing import List


class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        # dp_max[i]表示以nums[i]结尾最大的乘积
        # dp_min[i]表示以nums[i]结尾最小的乘积
        dp_max = [0] * len(nums)
        dp_min = [0] * len(nums)
        dp_max[0] = nums[0]
        dp_min[0] = nums[0]
        ans = nums[0]
        for i in range(1, len(nums)):
            if nums[i] < 0:
                dp_max[i] = max(nums[i], dp_min[i - 1] * nums[i])
                dp_min[i] = min(nums[i], dp_max[i - 1] * nums[i])
            else:
                dp_max[i] = max(nums[i], dp_max[i - 1] * nums[i])
                dp_min[i] = min(nums[i], nums[i] * dp_min[i - 1])
            ans = max(ans, dp_max[i])

        return ans
