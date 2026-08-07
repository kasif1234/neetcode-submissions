class Solution:
    def mergeKLists(self, lists: List[Optional[ListNode]]) -> Optional[ListNode]:
        if not lists:
            return None
        return self.divide(lists, 0, len(lists) - 1)

    def divide(self, lists, l, r):
        if l > r:
            return None
        if l == r:
            return lists[l]
        mid = (l + r) // 2
        left = self.divide(lists, l, mid)
        right = self.divide(lists, mid + 1, r)
        return self.mergeTwo(left, right)

    def mergeTwo(self, l1, l2):
        dummy = ListNode()
        tail = dummy
        while l1 and l2:
            if l1.val <= l2.val:
                tail.next = l1
                l1 = l1.next
            else:
                tail.next = l2
                l2 = l2.next
            tail = tail.next
        tail.next = l1 if l1 else l2
        return dummy.next