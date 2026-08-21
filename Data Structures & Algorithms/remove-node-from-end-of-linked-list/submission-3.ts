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
     * @param {ListNode} head
     * @param {number} n
     * @return {ListNode}
     */
    removeNthFromEnd(head: ListNode | null, n: number): ListNode {
        const dummy = new ListNode(0, head);
        let first: ListNode | null = dummy;
        let second: ListNode | null = dummy;

        for (let i = 0; i <= n; i++) {
            first = first!.next;
        }

        while (first !== null) {
            first = first.next;
            second = second!.next;
        }

        second!.next = second!.next!.next;

        return dummy.next as ListNode;
    }
}