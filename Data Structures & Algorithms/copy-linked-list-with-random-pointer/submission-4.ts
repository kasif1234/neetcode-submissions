// class Node {
//   constructor(val, next = null, random = null) {
//       this.val = val;
//       this.next = next;
//       this.random = random;
//   }
// }

class Solution {
    /**
     * @param {Node} head
     * @return {Node}
     */
    copyRandomList(head: Node | null): Node | null {
        if (!head) return null;

        const oldToCopy = new Map<Node, Node>();

        let cur: Node | null = head;
        while (cur) {
            oldToCopy.set(cur, new Node(cur.val));
            cur = cur.next;
        }

        cur = head;
        while (cur) {
            const copy = oldToCopy.get(cur)!;
            copy.next = cur.next ? oldToCopy.get(cur.next)! : null;
            copy.random = cur.random ? oldToCopy.get(cur.random)! : null;
            cur = cur.next;
        }

        return oldToCopy.get(head)!;
    }
}