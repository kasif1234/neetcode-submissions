class Solution {
    lowestCommonAncestor(
        root: TreeNode | null,
        p: TreeNode | null,
        q: TreeNode | null,
    ): TreeNode | null {
        let cur = root;
        while (cur) {
            if (p!.val < cur.val && q!.val < cur.val) {
                cur = cur.left;
            } else if (p!.val > cur.val && q!.val > cur.val) {
                cur = cur.right;
            } else {
                return cur;
            }
        }
        return null;
    }
}