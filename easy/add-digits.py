# https://leetcode.com/problems/add-digits/

class Solution:
    def addDigits(self, num: int) -> int:
        while num >= 10:
            numx = 0
            while num != 0:
                numx += num % 10
                num = int(num/10)
            num = numx
        return num


def main():
    foo = Solution()
    print(foo.addDigits(1))
    print(foo.addDigits(0))
    print(foo.addDigits(10))
    print(foo.addDigits(1231231))


if __name__ == "__main__":
    main()
