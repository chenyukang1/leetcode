# https://leetcode.cn/problems/valid-parentheses/description/?envType=study-plan-v2&envId=top-100-liked


import collections


class Solution:
    def isValid(self, s: str) -> bool:
        stack = collections.deque()
        for c in s:
            if c == "(" or c == "{" or c == "[":
                stack.append(c)
            else:
                if not stack:
                    return False
                top = stack.pop()
                if c == ")" and top != "(":
                    return False
                if c == "}" and top != "{":
                    return False
                if c == "]" and top != "[":
                    return False
        if stack:
            return False

        return True
