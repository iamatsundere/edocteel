# https://leetcode.com/problems/remove-duplicates-from-sorted-array/
from typing import List


class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        s = set()
        i=0
        while i < len(nums):
            # print(i,nums,len(nums))
            item = nums[i]
            if item in s:
                nums.remove(item)
            else:
                s.add(item)
                i+=1

        # print(nums)
        return len(s)


def main():
    foo = Solution()
    print(foo.removeDuplicates([1, 1, 2, 3]))
    print(foo.removeDuplicates([0,0,1,1,1,2,2,3,3,4]))


if __name__ == "__main__":
    main()
