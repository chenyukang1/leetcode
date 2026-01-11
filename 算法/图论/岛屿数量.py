# https://leetcode.cn/problems/number-of-islands/?envType=study-plan-v2&envId=top-100-liked

from collections import deque
from typing import List


class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        visited = [[False for _ in row] for row in grid]
        ans = 0
        for i, m in enumerate(grid):
            for j, n in enumerate(m):
                if visited[i][j]:
                    visited[i][j] = True
                    continue
                if n == '1':
                    bsf(i, j, grid, visited)
                    ans += 1
        return ans

directions = [[0, -1], [-1, 0], [0, 1], [1, 0]]

def bsf(i, j, grid, visited):
    queue = deque()
    queue.append((i, j))
    while len(queue) > 0:
        i,j=queue.popleft()
        for direction in directions:
            k, l = i + direction[0], j + direction[1]  # noqa: E741
            if k >= 0 and k < len(grid) and l >= 0 and l < len(grid[0]) and not visited[k][l] and grid[k][l] == '1':
                visited[k][l] = True
                queue.append((k,l))
