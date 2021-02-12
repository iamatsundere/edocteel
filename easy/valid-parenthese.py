# https://leetcode.com/problems/valid-parenthese/
from itertools import islice


class Solution:
    def isValid(self, s: str) -> bool:
        stack = []
        d = {
            '(': ')',
            '{': '}',
            '[': ']'
        }

        for item in islice(s, 0, len(s)):
            if item in d:
                stack.append(item)
            else:
                if len(stack) > 0:
                    itemx = stack.pop()
                    if itemx in d and d[itemx] == item:
                        continue
                    else:
                        stack.append(itemx)
                stack.append(item)
        if len(stack) == 0:
            return True
        else:
            return False


def main():
    foo = Solution()
    print(foo.isValid('('))
    print(foo.isValid(']'))
    print(foo.isValid('()'))
    print(foo.isValid('()[]{}'))
    print(foo.isValid('(]'))
    print(foo.isValid('{[]}'))
    print(foo.isValid('([)]'))


if __name__ == "__main__":
    main()
