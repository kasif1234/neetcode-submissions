class Solution {
    mergeKLists(lists: ListNode[]): ListNode {
        if (!lists || lists.length === 0) return null;
        return this.divide(lists, 0, lists.length - 1);
    }

    private divide(lists: ListNode[], l: number, r: number): ListNode {
        if (l > r) return null;
        if (l === r) return lists[l];
        const mid = Math.floor((l + r) / 2);
        const left = this.divide(lists, l, mid);
        const right = this.divide(lists, mid + 1, r);
        return this.mergeTwo(left, right);
    }

    private mergeTwo(l1: ListNode, l2: ListNode): ListNode {
        const dummy = new ListNode();
        let tail = dummy;
        while (l1 !== null && l2 !== null) {
            if (l1.val <= l2.val) {
                tail.next = l1;
                l1 = l1.next;
            } else {
                tail.next = l2;
                l2 = l2.next;
            }
            tail = tail.next;
        }
        tail.next = l1 !== null ? l1 : l2;
        return dummy.next;
    }
}