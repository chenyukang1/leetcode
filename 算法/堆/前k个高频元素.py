# https://leetcode.cn/problems/top-k-frequent-elements/?envType=study-plan-v2&envId=top-100-liked

import heapq
from typing import List


class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        count = {}
        for num in nums:
            if num not in count:
                count[num] = 1
            else:
                count[num] += 1

        min_heap = []

        for num in count.keys():
            if len(min_heap) < k:
                heapq.heappush(min_heap, (count[num], num))
            elif count[num] > min_heap[0][0]:
                heapq.heapreplace(min_heap, (count[num], num))

        return [item[1] for item in min_heap]


if __name__ == "__main__":
    solution = Solution()
    print(solution.topKFrequent([1, 1, 1, 2, 2, 3], 2))
