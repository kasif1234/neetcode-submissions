class TrieNode {
public:
    TrieNode* children[26] = {};
    string word = "";
};

class Solution {
public:
    vector<string> res;
    int ROWS, COLS;

    void insert(TrieNode* root, const string& word) {
        TrieNode* node = root;
        for (char ch : word) {
            int i = ch - 'a';
            if (!node->children[i]) node->children[i] = new TrieNode();
            node = node->children[i];
        }
        node->word = word;
    }

    void dfs(vector<vector<char>>& board, int r, int c, TrieNode* node) {
        int idx = board[r][c] - 'a';
        if (idx < 0 || !node->children[idx]) return;
        TrieNode* child = node->children[idx];

        if (!child->word.empty()) {
            res.push_back(child->word);
            child->word = "";
        }

        char ch = board[r][c];
        board[r][c] = '#';
        int dr[4] = {1,-1,0,0}, dc[4] = {0,0,1,-1};
        for (int i = 0; i < 4; i++) {
            int nr = r + dr[i], nc = c + dc[i];
            if (nr >= 0 && nr < ROWS && nc >= 0 && nc < COLS && board[nr][nc] != '#')
                dfs(board, nr, nc, child);
        }
        board[r][c] = ch;

        bool hasChild = false;
        for (int i = 0; i < 26; i++) if (child->children[i]) { hasChild = true; break; }
        if (!hasChild) {
            delete child;
            node->children[idx] = nullptr;
        }
    }

    vector<string> findWords(vector<vector<char>>& board, vector<string>& words) {
        TrieNode* root = new TrieNode();
        for (auto& w : words) insert(root, w);

        ROWS = board.size();
        COLS = board[0].size();

        for (int r = 0; r < ROWS; r++)
            for (int c = 0; c < COLS; c++)
                dfs(board, r, c, root);

        return res;
    }
};