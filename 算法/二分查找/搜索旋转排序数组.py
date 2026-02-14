# https://leetcode.cn/problems/search-in-rotated-sorted-array/?envType=study-plan-v2&envId=top-100-liked

from typing import List


class Solution:
    # 思路一: 先找到分界点，然后确认落在哪个有序区间，做一次二分。虽然能过，但不是O(logN)算法
    def search(self, nums: List[int], target: int) -> int:
        j = 0
        for i in range(1, len(nums)):
            if nums[i] < nums[i - 1]:
                j = i
                break
        # [0, j-1] [j, len(nums)-1]

        def binary_search(nums, left, right, target):
            while left <= right:
                mid = left + ((right - left) >> 1)
                if nums[mid] == target:
                    return mid
                elif nums[mid] < target:
                    left = mid + 1
                else:
                    right = mid - 1
            return -1

        if nums[0] <= target <= nums[j - 1]:
            return binary_search(nums, 0, j - 1, target)

        return binary_search(nums, j, len(nums) - 1, target)

    # 直接二分，一定有一个区间是有序区间，如果落在有序区间则二分，否则去另一个区间继续
    def search2(self, nums: List[int], target: int) -> int:
        n = len(nums)
        left, right = 0, n - 1
        while left <= right:
            mid = left + ((right - left) >> 1)
            if nums[mid] == target:
                return mid
            if nums[0] <= nums[mid]:
                if nums[0] <= target <= nums[mid]:
                    right = mid - 1
                else:
                    left = mid + 1
            else:
                if nums[mid] <= target <= nums[n - 1]:
                    left = mid + 1
                else:
                    right = mid - 1

        return -1
