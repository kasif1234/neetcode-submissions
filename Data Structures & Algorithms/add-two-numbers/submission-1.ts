/**
 * Definition for singly-linked list.
 * class ListNode {
 *     constructor(val = 0, next = null) {
 *         this.val = val;
 *         this.next = next;
 *     }
 * }
 */

class Solution {
    /**
     * @param {ListNode} l1
     * @param {ListNode} l2
     * @return {ListNode}
     */
    addTwoNumbers(l1: ListNode | null, l2: ListNode | null): ListNode {
        const dummy = new ListNode();
        let curr = dummy;
        let carry = 0;

        while (l1 !== null || l2 !== null || carry !== 0) {
            const v1 = l1 !== null ? l1.val : 0;
            const v2 = l2 !== null ? l2.val : 0;

            const total = v1 + v2 + carry;
            carry = Math.floor(total / 10);
            curr.next = new ListNode(total % 10);
            curr = curr.next;

            l1 = l1 !== null ? l1.next : null;
            l2 = l2 !== null ? l2.next : null;
        }

        return dummy.next as ListNode;
    }
}