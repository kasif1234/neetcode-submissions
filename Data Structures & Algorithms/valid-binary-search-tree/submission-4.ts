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
    isValidBST(root: TreeNode | null): boolean {
        const valid = (node: TreeNode | null, left: number, right: number): boolean => {
            if (!node) return true;
            if (!(node.val > left && node.val < right)) return false;
            return valid(node.left, left, node.val) && valid(node.right, node.val, right);
        };
        
        return valid(root, -Infinity, Infinity);
    }
}