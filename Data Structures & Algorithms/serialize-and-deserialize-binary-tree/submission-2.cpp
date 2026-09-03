class Codec {
public:
    // Encodes a tree to a single string.
    string serialize(TreeNode* root) {
        string res;
        dfsSerialize(root, res);
        return res;
    }

    // Decodes your encoded data to tree.
    TreeNode* deserialize(string data) {
        int pos = 0;
        return dfsDeserialize(data, pos);
    }

private:
    void dfsSerialize(TreeNode* node, string& res) {
        if (!node) {
            res += "N,";
            return;
        }
        res += to_string(node->val) + ",";
        dfsSerialize(node->left, res);
        dfsSerialize(node->right, res);
    }

    TreeNode* dfsDeserialize(const string& data, int& pos) {
        // extract next token
        int start = pos;
        while (data[pos] != ',') pos++;
        string token = data.substr(start, pos - start);
        pos++; // skip comma

        if (token == "N") {
            return nullptr;
        }
        TreeNode* node = new TreeNode(stoi(token));
        node->left = dfsDeserialize(data, pos);
        node->right = dfsDeserialize(data, pos);
        return node;
    }
};