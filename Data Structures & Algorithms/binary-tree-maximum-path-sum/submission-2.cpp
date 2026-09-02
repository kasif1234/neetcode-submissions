class Solution {
public:
    int best = INT_MIN;

    int maxPathSum(TreeNode* root) {
        dfs(root);
        return best;
    }

    int dfs(TreeNode* node) {
        if (!node) return 0;
        int left = max(dfs(node->left), 0);
        int right = max(dfs(node->right), 0);
        best = max(best, node->val + left + right);
        return node->val + max(left, right);
    }
};