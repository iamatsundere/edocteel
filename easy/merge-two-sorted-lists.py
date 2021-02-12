# https://leetcode.com/problems/merge-two-sorted-lists/
# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def mergeTwoLists(self, l1: ListNode, l2: ListNode) -> ListNode:

        if l1 == None:
            return l2
        elif l2 == None:
            return l1
        
        l = ListNode()
        lx = l
        while l1 != None and l2 != None:
            if l1.val <= l2.val:
                l.val = l1.val
                l1 = l1.next
            else:
                l.val = l2.val
                l2 = l2.next
            if l1 != None and l2 != None:
                l.next = ListNode()
                l = l.next

        if l1 != None:
            l.next = l1
        if l2 != None:
            l.next = l2

        return lx


def main():
    foo = Solution()
    # print(foo.removeElements([1,2,6,3,4,5,6]))


if __name__ == "__main__":
    main()
