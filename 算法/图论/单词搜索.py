# https://leetcode.cn/problems/word-search/description/?envType=study-plan-v2&envId=top-100-liked

from typing import List


class Solution:
    def exist(self, board: List[List[str]], word: str) -> bool:
        if not board:
            return False
        rows, cols = len(board), len(board[0])
        visited = [[False for _ in range(cols)] for _ in range(rows)]
        directions = [[-1, 0], [1, 0], [0, -1], [0, 1]]
        def search(i, j, idx) -> bool:
            if idx == len(word):
                return True
            if i < 0 or i >= rows or j < 0 or j >= cols or visited[i][j] or board[i][j] != word[idx]:
                return False

            visited[i][j] = True
            for dir in directions:
                newX = j + dir[1]
                newY = i + dir[0]
                if search(newY, newX, idx+1):
                    return True

            visited[i][j] = False
            return False

        for i, row in enumerate(board):
            for j, letter in enumerate(row):
                if letter == word[0] and search(i, j, 0):
                    return True
        return False
