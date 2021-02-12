# https://leetcode.com/problems/longest-common-prefix/
from typing import List


class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        i = 0
        n = len(strs)
        str = ""
        if n > 0:
            s0 = strs[0]
            if len(s0) > 0:
                sx = s0[0:i]

                while i <= len(s0):
                    # print(sx)
                    count = 0
                    for item in strs:
                        if item == "":
                            return ""
                        if sx == item[0:i]:
                            count += 1
                        else:
                            break

                    if count == n:
                        str = sx
                        i += 1
                        sx = s0[0:i]
                    else:
                        break

        return str


def main():
    foo = Solution()
    print(foo.longestCommonPrefix(["flower", "flow", "flight"]))
    print(foo.longestCommonPrefix([" ", "dog", "racecar", "car"]))
    print(foo.longestCommonPrefix([""]))
    print(foo.longestCommonPrefix(["a"]))
    print(foo.longestCommonPrefix(["c", "acc", "ccc"]))


if __name__ == "__main__":
    main()
