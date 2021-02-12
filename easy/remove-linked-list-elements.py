# https://leetcode.com/problems/remove-linked-list-elements/
# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def removeElements(self, head: ListNode, val: int) -> ListNode:
        if head == None:
            return None

        head1 = head
        head2 = head1.next
        head0 = head1
        while head1 != None and head2 != None:
            if head2.val == val:
                head2 = head2.next
                head1.next = head2
            else:
                head1 = head1.next
                head2 = head2.next

        while head0 != None and head0.val == val:
            head0 = head0.next

        return head0


def main():
    foo = Solution()
    # print(foo.removeElements([1,2,6,3,4,5,6]))


if __name__ == "__main__":
    main()
