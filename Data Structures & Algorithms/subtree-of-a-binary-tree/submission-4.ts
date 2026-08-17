class Solution {
    isSubtree(root: TreeNode | null, subRoot: TreeNode | null): boolean {
        if (!subRoot) return true;
        if (!root) return false;

        if (this.sameTree(root, subRoot)) return true;
        return this.isSubtree(root.left, subRoot) || this.isSubtree(root.right, subRoot);
    }

    private sameTree(p: TreeNode | null, q: TreeNode | null): boolean {
        if (!p && !q) return true;
        if (!p || !q || p.val !== q.val) return false;
        return this.sameTree(p.left, q.left) && this.sameTree(p.right, q.right);
    }
}