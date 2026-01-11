
from typing import List


class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        path = []
        ans = []
        def backTrack(left, right):
            if len(path) == 2 * n:
                ans.append("".join(path))
                return
            if left < n:
                path.append("(")
                backTrack(left + 1, right)
                path.pop()
            if right < left:
                path.append(")")
                backTrack(left, right + 1)
                path.pop()
        backTrack(0, 0)
        return ans



def main():
    solution = Solution()
    print(solution.generateParenthesis(3))

if __name__ == "__main__":
    main()
