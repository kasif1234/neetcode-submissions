class Solution {
public:
    int preIdx = 0;
    vector<int> preorder;
    unordered_map<int, int> inMap;

    TreeNode* buildTree(vector<int>& preorder, vector<int>& inorder) {
        this->preorder = preorder;
        for (int i = 0; i < inorder.size(); i++) {
            inMap[inorder[i]] = i;
        }
        return build(0, inorder.size() - 1);
    }

    TreeNode* build(int l, int r) {
        if (l > r) return nullptr;

        int rootVal = preorder[preIdx++];
        TreeNode* root = new TreeNode(rootVal);

        int mid = inMap[rootVal];
        root->left = build(l, mid - 1);
        root->right = build(mid + 1, r);

        return root;
    }
};