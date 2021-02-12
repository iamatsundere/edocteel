# https://leetcode.com/problems/remove-element/
from typing import List


class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        i = 0
        while i < len(nums):
            # print(i, nums, len(nums))
            if nums[i] == val:
                nums.remove(nums[i])
            else:
                i+=1
        return len(nums)


def main():
    foo = Solution()
    print(foo.removeElement([1, 1, 2, 3], 1))
    print(foo.removeElement([0, 1, 2, 2, 3, 0, 4, 2], 2))
    print(foo.removeElement([3, 2, 2, 3], 3))


if __name__ == "__main__":
    main()
