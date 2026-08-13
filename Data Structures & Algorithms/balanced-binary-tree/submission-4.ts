/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     constructor(val = 0, left = null, right = null) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */

class Solution {
    /**
     * @param {TreeNode} root
     * @return {boolean}
     */
    isBalanced(root: TreeNode | null): boolean {
        const dfs = (node: TreeNode | null): number => {
            if (node === null) return 0;

            const left = dfs(node.left);
            if (left === -1) return -1;

            const right = dfs(node.right);
            if (right === -1) return -1;

            if (Math.abs(left - right) > 1) return -1;

            return Math.max(left, right) + 1;
        };

        return dfs(root) !== -1;
    }
}