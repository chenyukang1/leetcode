# https://leetcode.cn/problems/search-a-2d-matrix/?envType=study-plan-v2&envId=top-100-liked

from typing import List


class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        low, high = 0, len(matrix)
        while low < high:
           mid = low + ((high - low) >> 1)
           if matrix[mid][0] == target:
               return True
           elif matrix[mid][0] < target:
               low = mid + 1
           else:
               high = mid
        i = low - 1
        low, high = 0, len(matrix[i])
        while low < high:
            mid = low + ((high - low) >> 1)
            if matrix[i][mid] == target:
               return True
            elif matrix[i][mid] < target:
               low = mid + 1
            else:
               high = mid

        return False

def main():
    Solution().searchMatrix([[1,3,5,7],[10,11,16,20],[23,30,34,60]], 3)

if __name__ == "__main__":
    main()
