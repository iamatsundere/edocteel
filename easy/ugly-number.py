# https://leetcode.com/problems/ugly-number/
class Solution:
    def isUgly(self, num: int) -> bool:
        arr = [2, 3, 5]
        i = 0

        if num < 1:
            return False

        while True:
            if num % arr[i] == 0:
                num = num//arr[i]
            else:
                i += 1
                if i == len(arr):
                    break

        if num == 1:
            return True

        return False


def main():
    foo = Solution()
    print(foo.isUgly(1))
    print(foo.isUgly(0))
    print(foo.isUgly(10))
    print(foo.isUgly(30))


if __name__ == "__main__":
    main()
