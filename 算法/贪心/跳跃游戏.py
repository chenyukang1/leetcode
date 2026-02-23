# https://leetcode.cn/problems/jump-game/?envType=study-plan-v2&envId=top-100-liked

from typing import List


class Solution:
    def canJump(self, nums: List[int]) -> bool:
        max_reach = 0
        for i, num in enumerate(nums):
            if i > max_reach:
                return False
            max_reach = max(max_reach, num)

            if max_reach >= len(nums) - 1:
                return True

        return False
