/*
// Definition for a Node.
class Node {
public:
    int val;
    Node* next;
    Node* random;
    
    Node(int _val) {
        val = _val;
        next = NULL;
        random = NULL;
    }
};
*/

class Solution {
public:
    Node* copyRandomList(Node* head) {
        if (!head) return nullptr;

        unordered_map<Node*, Node*> oldToCopy;

        Node* cur = head;
        while (cur) {
            oldToCopy[cur] = new Node(cur->val);
            cur = cur->next;
        }

        cur = head;
        while (cur) {
            Node* copy = oldToCopy[cur];
            copy->next = cur->next ? oldToCopy[cur->next] : nullptr;
            copy->random = cur->random ? oldToCopy[cur->random] : nullptr;
            cur = cur->next;
        }

        return oldToCopy[head];
    }
};