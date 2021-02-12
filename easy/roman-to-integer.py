# https://leetcode.com/problems/roman-to-integer/

class Solution:
    def romanToInt(self, s: str) -> int:
        d = {
            'I': 1,
            'V': 5,
            'X': 10,
            'L': 50,
            'C': 100,
            'D': 500,
            'M': 1000
        }

        x = 0
        i = len(s)-1
        while i >= 0:
            c = s[i]
            # print(i,c)
            if (i > 0):
                if d[c] > d[s[i-1]]:
                    x += d[c]-d[s[i-1]]
                    i -= 1
                else:
                    x += d[c]
            else:
                x += d[c]
            i -= 1

        return x


def main():
    foo = Solution()
    print(foo.romanToInt('IV'))
    print(foo.romanToInt('III'))
    print(foo.romanToInt('IX'))
    print(foo.romanToInt('XXIIV'))
    print(foo.romanToInt('MCMXCIV'))


if __name__ == "__main__":
    main()
