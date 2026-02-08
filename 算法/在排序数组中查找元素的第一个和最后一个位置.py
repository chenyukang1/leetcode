# https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/description/?envType=study-plan-v2&envId=top-100-liked

from typing import List


class Solution:
    # 如果整个数组都相同，最坏情况O(n)
    def searchRangev1(self, nums: List[int], target: int) -> List[int]:
        left, right = 0, len(nums)
        while left < right:
            mid = left + ((right - left) >> 1)
            if nums[mid] == target:
                i, j = mid, mid
                while j + 1 < len(nums) and nums[j + 1] == nums[j]:
                    j += 1
                while i - 1 >= 0 and nums[i - 1] == nums[i]:
                    i -= 1
                return [i, j]
            elif nums[mid] < target:
                left = mid + 1
            else:
                right = mid

        return [-1, -1]

    def searchRangev2(self, nums: List[int], target: int) -> List[int]:
        left = self.binay_search_first_greater(nums, target, True)
        right = self.binay_search_first_greater(nums, target, False)
        if left == -1:
            return [-1, -1]
        elif right == -1:
            return [left, left]
        return [left, right - 1]

    def binay_search_first_greater(self, nums: List[int], target: int, eq: bool) -> int:
        left, right, ans = 0, len(nums), -1
        while left < right:
            mid = left + ((right - left) >> 1)
            if nums[mid] == target:
                if eq:
                    right = mid
                    ans = mid
                else:
                    left = mid + 1
            elif nums[mid] < target:
                left = mid + 1
            else:
                right = mid
                if not eq:
                    ans = mid
        return ans
