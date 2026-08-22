/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    if head == nil {
        return nil
    }

    oldToCopy := make(map[*Node]*Node)

    cur := head
    for cur != nil {
        oldToCopy[cur] = &Node{Val: cur.Val}
        cur = cur.Next
    }

    cur = head
    for cur != nil {
        copy := oldToCopy[cur]
        if cur.Next != nil {
            copy.Next = oldToCopy[cur.Next]
        }
        if cur.Random != nil {
            copy.Random = oldToCopy[cur.Random]
        }
        cur = cur.Next
    }

    return oldToCopy[head]
}