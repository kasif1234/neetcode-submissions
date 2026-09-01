class Solution {
    reorderList(head: ListNode | null): void {
        let slow: ListNode = head!;
        let fast: ListNode | null = head!.next;
        while (fast !== null && fast.next !== null) {
            slow = slow.next!;
            fast = fast.next.next;
        }

        let second: ListNode | null = slow.next;
        slow.next = null;
        let prev: ListNode | null = null;
        while (second !== null) {
            const temp: ListNode | null = second.next;
            second.next = prev;
            prev = second;
            second = temp;
        }
        second = prev;

        let first: ListNode | null = head;
        while (second !== null) {
            const temp1: ListNode | null = first!.next;
            const temp2: ListNode | null = second.next;
            first!.next = second;
            second.next = temp1;
            first = temp1;
            second = temp2;
        }
    }
}