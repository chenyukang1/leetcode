# https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/description/?envType=study-plan-v2&envId=top-100-liked# 直接二分，一定有一个区间是有序区间，如果落在有序区间则二分，否则去另一个区间继续

from typing import List


class Solution:
    def findMin(self, nums: List[int]) -> int:
        left, right = 0, len(nums) - 1
        while left < right:
            mid = left + ((right - left) >> 1)
            if nums[mid] < nums[len(nums) - 1]:
                right = mid
            else:
                left = mid + 1

        return nums[left]
