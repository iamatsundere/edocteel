# https://leetcode.com/problems/count-primes/
class Solution:
    def countPrimes(self, n: int) -> int:

        if n <= 1:
            return 0

        A = [True]*(n+1)
        A[0] = False
        A[1] = False
        x = n-1
        for i in range(2, n//2+1):
            for j in range(2, n//2+1):
                # print(i, j, i*j)
                if i*j > n:
                    break
                if A[i*j] == True:
                    A[i*j] = False
                    x -= 1

        if A[n] == True:
            x -= 1

        return x


def main():
    foo = Solution()
    print(foo.countPrimes(10))
    print(foo.countPrimes(0))
    print(foo.countPrimes(2))
    # print(foo.countPrimes(30))


if __name__ == "__main__":
    main()
