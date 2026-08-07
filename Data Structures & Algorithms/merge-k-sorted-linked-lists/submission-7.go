/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }
    return divide(lists, 0, len(lists)-1)
}

func divide(lists []*ListNode, l, r int) *ListNode {
    if l > r {
        return nil
    }
    if l == r {
        return lists[l]
    }
    mid := l + (r-l)/2
    left := divide(lists, l, mid)
    right := divide(lists, mid+1, r)
    return mergeTwo(left, right)
}

func mergeTwo(l1, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    tail := dummy
    for l1 != nil && l2 != nil {
        if l1.Val <= l2.Val {
            tail.Next = l1
            l1 = l1.Next
        } else {
            tail.Next = l2
            l2 = l2.Next
        }
        tail = tail.Next
    }
    if l1 != nil {
        tail.Next = l1
    } else {
        tail.Next = l2
    }
    return dummy.Next
}