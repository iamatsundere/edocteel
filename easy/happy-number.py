# https://leetcode.com/problems/happy-number/

class Solution:
    def isHappy(self, n: int) -> bool:

        if n == 0:
            return False

        s=set()
        while n not in s:
            s.add(n)
            # print(s)
            nx = 0
            while n != 0:
                a = n % 10
                nx += a**2
                n = n // 10
                # print(a, nx, n)
            n = nx
            # print(n)
            if n==1:
                return True

        return False


def main():
    foo = Solution()
    print(foo.isHappy(19))
    print(foo.isHappy(1))
    print(foo.isHappy(0))
    print(foo.isHappy(1000))
    print(foo.isHappy(123))


if __name__ == "__main__":
    main()
