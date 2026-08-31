class Solution {
    reverseKGroup(head: ListNode | null, k: number): ListNode | null {
        const dummy = new ListNode(0, head);
        let groupPrev: ListNode = dummy;

        while (true) {
            let kth: ListNode | null = groupPrev;
            for (let i = 0; i < k && kth !== null; i++) {
                kth = kth.next;
            }
            if (kth === null) break;

            const groupNext: ListNode | null = kth.next;
            let prev: ListNode | null = groupNext;
            let curr: ListNode | null = groupPrev.next;

            while (curr !== groupNext) {
                const temp: ListNode | null = curr!.next;
                curr!.next = prev;
                prev = curr;
                curr = temp;
            }

            const temp = groupPrev.next;
            groupPrev.next = kth;
            groupPrev = temp!;
        }

        return dummy.next;
    }
}