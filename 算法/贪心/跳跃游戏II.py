# https://leetcode.cn/problems/jump-game-ii/?envType=study-plan-v2&envId=top-100-liked

from typing import List


class Solution:
    def jump(self, nums: List[int]) -> int:
        left, right = 0, 0
        max_reach = 0
        count = 0
        while max_reach < len(nums) - 1:
            for i in range(left, right + 1):
                max_reach = max(max_reach, i + nums[i])
            left = right + 1
            right = max_reach
            count += 1
        return count
