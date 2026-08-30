class Solution {
    buildTree(preorder: number[], inorder: number[]): TreeNode {
        const inMap = new Map<number, number>();
        for (let i = 0; i < inorder.length; i++) {
            inMap.set(inorder[i], i);
        }
        let preIdx = 0;

        const build = (l: number, r: number): TreeNode | null => {
            if (l > r) return null;

            const rootVal = preorder[preIdx++];
            const root = new TreeNode(rootVal);

            const mid = inMap.get(rootVal)!;
            root.left = build(l, mid - 1);
            root.right = build(mid + 1, r);

            return root;
        };

        return build(0, inorder.length - 1);
    }
}