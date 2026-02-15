# https://leetcode.cn/problems/decode-string/?envType=study-plan-v2&envId=top-100-liked


class Solution:
    def decodeString(self, s: str) -> str:
        stack = []
        cur_str = ""
        cur_num = 0
        for char in s:
            if char.isdigit():
                cur_num = cur_num * 10 + int(char)
            elif char == "[":
                stack.append((cur_str, cur_num))
                cur_str = ""
                cur_num = 0
            elif char == "]":
                last_str, num = stack.pop()
                cur_str = last_str + num * cur_str
            else:
                cur_str += char

        return cur_str


if __name__ == "__main__":
    solution = Solution()
    ans = solution.decodeString("3[a]2[bc]")
    print(ans)
