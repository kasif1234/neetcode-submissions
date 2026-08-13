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
     * @return {number}
     */
    diameterOfBinaryTree(root: TreeNode | null): number {
        let diameter = 0;

        const height = (node: TreeNode | null): number => {
            if (node === null) return 0;

            const left = height(node.left);
            const right = height(node.right);

            diameter = Math.max(diameter, left + right);

            return Math.max(left, right) + 1;
        };

        height(root);
        return diameter;
    }
}