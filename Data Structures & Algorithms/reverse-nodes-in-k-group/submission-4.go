func reverseKGroup(head *ListNode, k int) *ListNode {
    dummy := &ListNode{Next: head}
    groupPrev := dummy

    for {
        kth := groupPrev
        i := 0
        for i < k && kth != nil {
            kth = kth.Next
            i++
        }
        if kth == nil {
            break
        }

        groupNext := kth.Next
        prev, curr := groupNext, groupPrev.Next
        for curr != groupNext {
            temp := curr.Next
            curr.Next = prev
            prev = curr
            curr = temp
        }

        temp := groupPrev.Next
        groupPrev.Next = kth
        groupPrev = temp
    }

    return dummy.Next
}