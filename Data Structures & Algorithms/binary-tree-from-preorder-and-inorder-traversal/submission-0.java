class Solution {
    private int preIdx = 0;
    private int[] preorder;
    private Map<Integer, Integer> inMap = new HashMap<>();

    public TreeNode buildTree(int[] preorder, int[] inorder) {
        this.preorder = preorder;
        for (int i = 0; i < inorder.length; i++) {
            inMap.put(inorder[i], i);
        }
        return build(0, inorder.length - 1);
    }

    private TreeNode build(int l, int r) {
        if (l > r) return null;

        int rootVal = preorder[preIdx++];
        TreeNode root = new TreeNode(rootVal);

        int mid = inMap.get(rootVal);
        root.left = build(l, mid - 1);
        root.right = build(mid + 1, r);

        return root;
    }
}