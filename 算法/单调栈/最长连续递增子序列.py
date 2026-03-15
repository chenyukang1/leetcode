from typing import List


class Solution:
    def findLengthOfLCIS(self, nums: List[int]) -> int:
        ans = 1
        cur_len = 1
        for r in range(1, len(nums)):
            if nums[r] > nums[r - 1]:
                cur_len += 1
            else:
                cur_len = 1
            ans = max(ans, cur_len)

        return ans
