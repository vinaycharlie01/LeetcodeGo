
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def getVal(self, l1: Optional[ListNode]):
        if l1 == None:
            return 0
        return l1.Val
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        a=ListNode
        head = a
        carry, val = 0, 0
        while l1 != None or l2 != None or carry != 0:
            val = carry + self.getVal(l1) + self.getVal(l2)
            carry = val // 10
            val = val % 10
            NewNode =ListNode(val)
            head.Next = NewNode
            head = head.Next
            if l1 != None:
                l1 = l1.Next
            if l2 != None:
                l2 = l2.Next
        return a.Next
