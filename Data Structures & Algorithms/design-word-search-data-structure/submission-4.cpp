class TrieNode {
public:
    TrieNode* children[26] = {nullptr};
    bool isEnd = false;
};

class WordDictionary {
private:
    TrieNode* root;

    bool dfs(TrieNode* node, const string& word, int i) {
        if (node == nullptr) return false;
        if (i == (int)word.size()) {
            return node->isEnd;
        }
        char c = word[i];
        if (c == '.') {
            for (int j = 0; j < 26; j++) {
                if (node->children[j] != nullptr && dfs(node->children[j], word, i + 1)) {
                    return true;
                }
            }
            return false;
        } else {
            return dfs(node->children[c - 'a'], word, i + 1);
        }
    }

public:
    WordDictionary() {
        root = new TrieNode();
    }

    void addWord(string word) {
        TrieNode* cur = root;
        for (char c : word) {
            int idx = c - 'a';
            if (cur->children[idx] == nullptr) {
                cur->children[idx] = new TrieNode();
            }
            cur = cur->children[idx];
        }
        cur->isEnd = true;
    }

    bool search(string word) {
        return dfs(root, word, 0);
    }
};

/**
 * Your WordDictionary object will be instantiated and called as such:
 * WordDictionary* obj = new WordDictionary();
 * obj->addWord(word);
 * bool param_2 = obj->search(word);
 */