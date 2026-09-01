func reorderList(head *ListNode) {
    slow, fast := head, head.Next
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    second := slow.Next
    slow.Next = nil
    var prev *ListNode
    for second != nil {
        temp := second.Next
        second.Next = prev
        prev = second
        second = temp
    }
    second = prev

    first := head
    for second != nil {
        temp1, temp2 := first.Next, second.Next
        first.Next = second
        second.Next = temp1
        first = temp1
        second = temp2
    }
}