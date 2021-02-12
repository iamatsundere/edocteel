# https://leetcode.com/problems/move-zeroes/
from typing import List


class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        count = 0
        i = 0
        while i < len(nums):
            # print(i, nums[i], nums)
            if nums[i] == 0:
                count += 1
                nums.pop(i)
            else:
                i += 1
        # print(nums)

        while count > 0:
            nums.append(0)
            count -= 1

        # print(nums)


def main():
    foo = Solution()
    print(foo.moveZeroes([0, 0, 1]))
    # print(foo.moveZeroes([1, 1, 2, 3]))
    # print(foo.moveZeroes([0, 1, 2, 2, 3, 0, 4, 2]))
    # print(foo.moveZeroes([3, 2, 2, 3]))


if __name__ == "__main__":
    main()
